package app

import (
	"log"
	"net/http"
	"time"

	"github.com/MustaphaSakka/traney/domain"
	"github.com/MustaphaSakka/traney/logger"
	"github.com/MustaphaSakka/traney/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	router := mux.NewRouter()
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewClientRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := ClientHandlers{service.NewClientService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// define routes
	router.HandleFunc("/clients", ch.getAllClients).Methods(http.MethodGet)
	router.HandleFunc("/clients/{client_id:[0-9]+}", ch.getClient).Methods(http.MethodGet)
	router.HandleFunc("/clients/{client_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:@/traney")
	if err != nil {
		panic(err)
	}

	logger.Info("Connexion to database is established.")

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
