package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Connection struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

var connections []Connection

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

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
	json.NewEncoder(w).Encode(&Connection{})
}

func createConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var connection Connection

	_ = json.NewDecoder(r.Body).Decode(&connection)
	connection.ID = strconv.Itoa(rand.Intn(1000000))
	connections = append(connections, connection)
	json.NewEncoder(w).Encode(connection)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	r.HandleFunc("/connections", getConnections).Methods("GET")
	r.HandleFunc("/connections/{id}", getConnection).Methods("GET")
	r.HandleFunc("/connections", createConnection).Methods("POST")

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8005", handler))
}
