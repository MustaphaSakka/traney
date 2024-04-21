package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MustaphaSakka/traney/service"
)

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
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, client)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
