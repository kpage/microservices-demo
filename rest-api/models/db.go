package models

import (
	"database/sql"
	"fmt"
	"os"

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
	// TODO: print helpful error message when required env vars are not passed
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
