package models

import (
	"database/sql"

	// Necessary to import mysql drivers
	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
	AllBooks() ([]*Book, error)
}

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	// TODO: pull these out into env vars?  How does this work best with docker-compose and kubernetes?
	//  - username
	//  - password
	//  - db hostname
	//  - db port
	//  - db name
	db, err := sql.Open("mysql", "root:rewt@tcp(db:3306)/restbucks")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
