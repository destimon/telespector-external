package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Connection struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var connections []Connection

func getConnection(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(connections)
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

	r.HandleFunc("/connections", getConnection).Methods("GET")
	r.HandleFunc("/connections", createConnection).Methods("POST")

	log.Fatal(http.ListenAndServe(":8005", r))
}
