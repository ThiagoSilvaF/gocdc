package sqlserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	kafka "github.com/gocdc/kafka"
)

var db *sql.DB

func InitDB(connector Connector) {
	var connString string
	if connector.DbName == "" {
		connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;",
			connector.DbHost, connector.DbUser, connector.DbPass, connector.DbPort)
	} else {
		connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
			connector.DbHost, connector.DbUser, connector.DbPass, connector.DbPort, connector.DbName)
	}

	var errdb error
	db, errdb = sql.Open("mssql", connString)
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}

	err := db.Ping()
	if err != nil {
		panic("PANIC when pinging db: " + err.Error())
	}

	go func() {
		for true {
			executeCdc()
		}
	}()
	return
}

func executeCdc() {
	for _, v := range Connectors {
		var query string
		if v.DbName != "" {
			query = fmt.Sprintf("SELECT * FROM %s.cdc.%s_%s_ct;",
				v.DbName, v.DbSchema, v.TableName)
		} else {
			query = fmt.Sprintf("SELECT * FROM cdc.%s_%s_ct;",
				v.DbSchema, v.TableName)
		}

		payload, err := getJSON(query)
		if err != nil {
			log.Fatal(err.Error())
		}
		kafka.SendMessage(v.KafkaBrokers, v.KafkaTopic, payload)
	}
}

func getJSON(sqlString string) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		fmt.Println(err.Error())

		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
