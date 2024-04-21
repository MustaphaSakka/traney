package domain

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ClientRepositoryDb struct {
	client *sqlx.DB
}

func (d ClientRepositoryDb) FindAll() ([]Client, error) {
	clients := make([]Client, 0)
	var err error

	findAllSql := "select * from clients"

	err = d.client.Select(&clients, findAllSql)
	if err != nil {
		log.Println("Error while querying `clients` table " + err.Error())
		return nil, err
	}

	return clients, nil
}

func NewClientRepositoryDb() ClientRepositoryDb {
	client, err := sqlx.Open("mysql", "root:@/traney")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connexion de database is established.")

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return ClientRepositoryDb{client}
}
