package main 

import (
    "fmt"
  
    _ "github.com/lib/pq"
    "github.com/ThiagoSilvaF/gocdc/postgres"
    "github.com/ThiagoSilvaF/gocdc/kafka"

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
