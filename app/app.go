package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/start", start).Methods(http.MethodGet)
	router.HandleFunc("/clients", getAllClients).Methods(http.MethodGet)
	router.HandleFunc("/clients/{client_id:[0-9]+}", getClient).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["client_id"])
}
