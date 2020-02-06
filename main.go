package main 

import (
    "fmt"
  
    _ "github.com/lib/pq"
    "github.com/ThiagoSilvaF/gocdc/utility/postgres"
    "github.com/ThiagoSilvaF/gocdc/utility/kafka"
	"github.com/ThiagoSilvaF/gocdc/utility/api"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "inventory"
)

func main() {

	fmt.Printf("going to call kafka")

	api.InitAPI()
	kafka.SendMessage()

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s sslmode=disable",
      host, port, user, password)

    postgres.InitDB(psqlInfo)
  

}
