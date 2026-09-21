package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MustaphaSakka/traney-lib/logger"
	"github.com/MustaphaSakka/traney/domain"
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
	router.HandleFunc("/clients", ch.getAllClients).Methods(http.MethodGet).Name("GetAllClients")
	router.HandleFunc("/clients/{client_id:[0-9]+}", ch.getClient).Methods(http.MethodGet).Name("GetClient")
	router.HandleFunc("/clients/{client_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost).Name("NewAccount")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getDbClient() *sqlx.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, database,
	)

	client, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	logger.Info("Connexion to database is established.")

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
