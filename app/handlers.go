package app

import (
	"encoding/json"
	"net/http"

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
