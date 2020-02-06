package main 

import (
    "fmt"
  
    _ "github.com/lib/pq"
<<<<<<< HEAD
    "gocdc/postgres"
    "gocdc/kafka"
=======
    "github.com/ThiagoSilvaF/gocdc/postgres"
    "github.com/ThiagoSilvaF/gocdc/kafka"
>>>>>>> e315284927f0f61ecd881353b0e80297b78869b1

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

  fmt.Printf("going to call kafka")

  kafka.SendMessage()

  if db_name == "POSTGRES" {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s sslmode=disable",
      host, port, user, password)

    postgres.InitDB(psqlInfo)
  }

}
