package main

import (
	"log"
	"net/http"

	"api-articles/pkg/api"
	"api-articles/pkg/db"
	"api-articles/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, _ := dbConn.DB()
		sqlDB.Close()
	}()

	// Create router
	router := mux.NewRouter()
	api.ArticlesRoutes(router, dbConn)

	// Start server
	log.Fatal(http.ListenAndServe(":2222", middleware.JSONContentType(router)))
}
