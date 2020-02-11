package main

import (
	"github.com/gocdc/api"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "inventory"
)

func main() {
	log.Info("*** Initializing APP ***")
	api.InitAPI()
}
