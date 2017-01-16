package models

import (
	"database/sql"
	"fmt"
	"os"

	// Necessary to import mysql drivers
	_ "github.com/go-sql-driver/mysql"
	log "github.com/mgutz/logxi/v1"
)

type Datastore interface {
	AllOrders() ([]*Order, error)
}

type DB struct {
	*sql.DB
}

func NewDB(logger log.Logger) (*DB, error) {
	// TODO: print helpful error message when required env vars are not passed
	logger.Info("Opening database...")
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
