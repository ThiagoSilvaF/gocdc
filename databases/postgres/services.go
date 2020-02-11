package postgres

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// Connector struct
type Connector struct {
	ConnectorName string   `json:"connector_name"`
	DbHost        string   `json:"db_host"`
	DbPort        int      `json:"db_port"`
	DbUser        string   `json:"db_user"`
	DbPass        string   `json:"db_pass"`
	DbName        string   `json:"db_name"`
	DbSlot        string   `json:"db_slot"`
	KafkaBrokers  []string `json:"kafka_brokers"`
	KafkaTopic    string   `json:"kafka_topic"`
}

var Connectors = make(map[string]Connector)

func GetConnectors(w http.ResponseWriter, r *http.Request) {
	log.Info("Calling getConnectors")
	createHttpReturn(w, http.StatusOK, Connectors, nil)
	return
}

func CreateConnector(w http.ResponseWriter, r *http.Request) {
	log.Info("Calling createConnector")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		createHttpReturn(w, http.StatusBadRequest, nil, err)
		return
	}

	var newConnector Connector
	if err := json.Unmarshal(body, &newConnector); err != nil {
		createHttpReturn(w, http.StatusBadRequest, nil, err)
		return
	}

	Connectors[newConnector.ConnectorName] = newConnector

	InitDB(newConnector)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newConnector)
	return
}

func GetConnector(w http.ResponseWriter, r *http.Request) {
	log.Info("Calling getConnector")
	params := mux.Vars(r)
	if params["connector_name"] == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Connectors[params["connector_name"]])
	return
}

func UpdateConnector(w http.ResponseWriter, r *http.Request) {
	log.Info("Calling updateConnector")

	createHttpReturn(w, http.StatusOK, Connectors, nil)
	return
}

func DeleteConnector(w http.ResponseWriter, r *http.Request) {
	log.Info("Calling deleteConnector")

	createHttpReturn(w, http.StatusOK, Connectors, nil)
	return
}

func createHttpReturn(w http.ResponseWriter, httpStatus int, Connectors map[string]Connector, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		json.NewEncoder(w).Encode(Connectors)
	}
}