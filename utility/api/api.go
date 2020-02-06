package api

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)

// Connector struct
type Connector struct {
	ID string `json:"id"`
  	Title string `json:"title"`
  	Body string `json:"body"` 
}

var connectors []Connector

// InitAPI - Initialize the Connector's endpoints
func InitAPI() {
	router := mux.NewRouter()  
	
	router.HandleFunc("/connectors", getConnectors).Methods("GET")
	router.HandleFunc("/connectors", createConnector).Methods("POST")
	router.HandleFunc("/connectors/{id}", getConnector).Methods("GET")
	router.HandleFunc("/connectors/{id}", updateConnector).Methods("PUT")
	router.HandleFunc("/connectors/{id}", deleteConnector).Methods("DELETE")
	
	http.ListenAndServe(":8000", router)
}

func getConnectors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(connectors)	
}

func createConnector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(connectors)	
}

func getConnector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(connectors)	
}

func updateConnector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(connectors)	
}

func deleteConnector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  	json.NewEncoder(w).Encode(connectors)	
}