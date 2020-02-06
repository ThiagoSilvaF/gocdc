package postgres

import (
    "database/sql"
    _ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
    "time"
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

//InitDB - initialize connection with PostgresDB
func InitDB(dataSourceName string)   {
	log.Info("Initializing connection with PostgresDB!")

    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
   }
   
   slots := postgresqlSlots{}
   for true {
    cdc(&slots, db)
    time.Sleep(2 * time.Second)

    log.Info(slots)
  } 
	
}

func cdc(slots *postgresqlSlots, db *sql.DB) {

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
      
    }

    err = rows.Err()
    if err != nil {
      log.Panic(err)
    }


}