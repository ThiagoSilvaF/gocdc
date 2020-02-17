package api

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	pg "github.com/gocdc/databases/postgres"
	sqlserver "github.com/gocdc/databases/sqlserver"
)

// InitAPI - Initialize the Connector's endpoints
func InitAPI() {
	log.Info("Initializing Rest API")
	router := mux.NewRouter()

	router.HandleFunc("/connectors/postgres", pg.GetConnectors).Methods("GET")
	router.HandleFunc("/connectors/postgres", pg.CreateConnector).Methods("POST")
	router.HandleFunc("/connectors/postgres/{id}", pg.GetConnector).Methods("GET")
	router.HandleFunc("/connectors/postgres/{id}", pg.UpdateConnector).Methods("PUT")
	router.HandleFunc("/connectors/postgres/{id}", pg.DeleteConnector).Methods("DELETE")

	router.HandleFunc("/connectors/sqlserver", sqlserver.GetConnectors).Methods("GET")
	router.HandleFunc("/connectors/sqlserver", sqlserver.CreateConnector).Methods("POST")
	router.HandleFunc("/connectors/sqlserver/{id}", sqlserver.GetConnector).Methods("GET")
	router.HandleFunc("/connectors/sqlserver/{id}", sqlserver.UpdateConnector).Methods("PUT")
	router.HandleFunc("/connectors/sqlserver/{id}", sqlserver.DeleteConnector).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}