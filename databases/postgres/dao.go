package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	kafka "github.com/gocdc/kafka"
)

var db *sql.DB

type postgresqlSlot struct {
	lsn  string
	xid  string
	data string
}

type postgresqlSlots struct {
	PostgresqlSlots []postgresqlSlot
}

//initDB - initialize connection with PostgresDB
func InitDB(connector Connector) {
	log.Info("Initializing connection with PostgresDB!")

	var pgInfo string
	if connector.DbName == "" {
		pgInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s sslmode=disable",
			connector.DbHost, connector.DbPort, connector.DbUser, connector.DbPass)
	} else {
		pgInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			connector.DbHost, connector.DbPort, connector.DbUser, connector.DbPass, connector.DbName)
	}

	var err error
	db, err = sql.Open("postgres", pgInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	go func() {
		for true {
			executeCdc(connector)
		}
	}()
	return
}

func executeCdc(conn Connector) error {
	slots := postgresqlSlots{}

	rows, err := db.Query("SELECT * FROM pg_logical_slot_get_changes('slot', null, null);")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		slot := postgresqlSlot{}
		err = rows.Scan(
			&slot.lsn,
			&slot.xid,
			&slot.data,
		)
		if err != nil {
			log.Panic(err)
		}
		slots.PostgresqlSlots = append(slots.PostgresqlSlots, slot)
		payload := slot.lsn + ", " + slot.xid + ", " + slot.data
		kafka.SendMessage(conn.KafkaBrokers, conn.KafkaTopic, payload)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}