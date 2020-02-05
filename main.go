package main 

import (
    "database/sql"
    "fmt"
  
    _ "github.com/lib/pq"
    "github.com/gocdc/postgres"
    "github.com/gocdc/kafka"

)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "inventory"
)

var db_name = "POSTGRES"

func main() {

  if db_name == "POSTGRES" {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s sslmode=disable",
      host, port, user, password)

    postgres.InitDB(psqlInfo)
  }

  kafka.SendMessage()
}
