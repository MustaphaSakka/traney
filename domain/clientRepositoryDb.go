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

func (d ClientRepositoryDb) FindById(id string) (*Client, error) {
	clientSql := "select client_id, name, city, zipcode, date_of_birth, status from clients where client_id = ?"
	var c Client

	err := d.client.Get(&c, clientSql, id)

	if err != nil {
		log.Println("Error while scanning client " + err.Error())
		return nil, err
	}
	return &c, nil
}

func NewClientRepositoryDb() ClientRepositoryDb {
	client, err := sqlx.Open("mysql", "root:@/traney")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connexion to database is established.")

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return ClientRepositoryDb{client}
}
