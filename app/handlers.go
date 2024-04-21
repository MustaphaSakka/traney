package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MustaphaSakka/traney/service"
)

// DTO
type Client struct {
	Name string `json:"full_name"`
	City string `json:"city"`
}

type ClientHandlers struct {
	service service.ClientService
}

func (ch *ClientHandlers) getAllClients(w http.ResponseWriter, r *http.Request) {
	clients, _ := ch.service.GetAllClient()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func (ch *ClientHandlers) getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["client_id"]

	client, err := ch.service.GetClient(id)
	fmt.Print(err)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(client)
	}
}
