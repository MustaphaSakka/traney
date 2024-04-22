package domain

import (
	"database/sql"

	"github.com/MustaphaSakka/traney/exception"
	"github.com/MustaphaSakka/traney/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ClientRepositoryDb struct {
	client *sqlx.DB
}

func (d ClientRepositoryDb) FindAll() ([]Client, *exception.AppException) {
	clients := make([]Client, 0)
	var err error

	findAllSql := "select * from clients"

	err = d.client.Select(&clients, findAllSql)
	if err != nil {
		logger.Error("Exception while querying `clients` table " + err.Error())
		return nil, exception.InternalServerException("Exception while querying database")
	}

	return clients, nil
}

func (d ClientRepositoryDb) FindById(id string) (*Client, *exception.AppException) {
	clientSql := "select client_id, name, city, zipcode, date_of_birth, status from clients where client_id = ?"
	var c Client

	err := d.client.Get(&c, clientSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NotFoundException("Client not found")
		} else {
			logger.Error("Exception while scanning client table" + err.Error())
			return nil, exception.InternalServerException("Unexpected database error")
		}
	}
	return &c, nil
}

func NewClientRepositoryDb(dbClient *sqlx.DB) ClientRepositoryDb {
	return ClientRepositoryDb{dbClient}
}
