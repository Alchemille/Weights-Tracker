package main

import (
	"fmt"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func AutoMigration(db *gorm.DB) {

	db.AutoMigrate(&Weight{})
}

func defaultRouteHdl(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte("Did you mean to contact the /weights route?"))
	return
}

func handleRoutes(db *gorm.DB, mux *http.ServeMux) {

	svc := WeightService{db}
	mux.HandleFunc("/weights", svc.handleWeights)
	mux.HandleFunc("/", defaultRouteHdl)

}

func main() {
	// Load database
	dsn := "host=localhost dbname=weights"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	AutoMigration(db)

	// Define the HTTP request multiplexer
	mux := http.NewServeMux()
	handleRoutes(db, mux)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

}
