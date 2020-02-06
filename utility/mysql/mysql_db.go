package models

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)


func InitDB(dataSourceName string) *sql.DB {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }
}

func cdc(slots *postgresqlSlots, db *sql.DB) (string, error) {

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
        return err
      }

      slots.PostgresqlSlots = append(slots.PostgresqlSlots, slot)
      
    }
    err = rows.Err()
    if err != nil {
      return err
    }
    return nil

}