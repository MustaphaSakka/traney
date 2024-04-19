package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DTO
type Client struct {
	Name string `json:"full_name"`
	City string `json:"city"`
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Let's start and enjoy this GO project!")
}

func getAllClients(w http.ResponseWriter, r *http.Request) {
	clients := []Client{
		{"Pierre", "Lyon"},
		{"Samuel", "Paris"},
		{"Isabelle", "Nice"},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}
