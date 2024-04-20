package app

import (
	"log"
	"net/http"

	"github.com/MustaphaSakka/traney/domain"
	"github.com/MustaphaSakka/traney/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	ch := ClientHandlers{service.NewClientService(domain.NewClientRepositoryDb())}

	// define routes
	router.HandleFunc("/clients", ch.getAllClients).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
