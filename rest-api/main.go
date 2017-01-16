package main

import (
	"microservices-demo/rest-api/api"
	"microservices-demo/rest-api/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mgutz/logxi/v1"
)

func main() {
	logger := log.New("")
	logger.Info("rest-api starting...")
	db, err := models.NewDB(logger)
	if err != nil {
		logger.Fatal("Could not open database", "err", err)
	}
	logger.Info("rest-api is up and serving on 3001")
	logger.Fatal("Serving", "err", http.ListenAndServe(":3001", api.Handlers(logger, db)))
}
