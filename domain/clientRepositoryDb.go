package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ClientRepositoryDb struct {
	client *sql.DB
}

func (d ClientRepositoryDb) FindAll() ([]Client, error) {
	findAllSql := "select * from clients"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying `clients` table " + err.Error())
		return nil, err
	}

	clients := make([]Client, 0)
	for rows.Next() {
		var c Client
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning `clients` table " + err.Error())
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func NewClientRepositoryDb() ClientRepositoryDb {
	client, err := sql.Open("mysql", "root:@/traney")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connexion de database is established.")

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return ClientRepositoryDb{client}
}
