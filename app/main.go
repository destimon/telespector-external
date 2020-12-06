package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Connection struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

var connections []Connection

func getConnections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(connections)
}

func getConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range connections {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(nil)
}

func createConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var connection Connection

	_ = json.NewDecoder(r.Body).Decode(&connection)
	connections = append(connections, connection)
	json.NewEncoder(w).Encode(connection)
}

func updateConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range connections {
		if item.ID == params["id"] {
			connections = append(connections[:index], connections[index+1:]...)
			var connection Connection
			_ = json.NewDecoder(r.Body).Decode(&connection)
			connection.ID = params["id"]
			connections = append(connections, connection)
			json.NewEncoder(w).Encode(connection)
			return
		}
	}
	json.NewEncoder(w).Encode(connections)
}

func deleteConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range connections {
		if item.ID == params["id"] {
			connections = append(connections[:index], connections[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(connections)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	r.HandleFunc("/connections/{id}", deleteConnection).Methods("DELETE")
	r.HandleFunc("/connections/{id}", updateConnection).Methods("PUT")
	r.HandleFunc("/connections", getConnections).Methods("GET")
	r.HandleFunc("/connections/{id}", getConnection).Methods("GET")
	r.HandleFunc("/connections", createConnection).Methods("POST")

	log.Print(http.ListenAndServe(":8005", handler))
}
