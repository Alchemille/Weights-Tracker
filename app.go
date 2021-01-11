package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

type Weight struct {
	gorm.Model
	Value int       `json:"value"`
	Date  time.Time `json:"date"`
}

type WeightService struct {
	db *gorm.DB
}

func (svc *WeightService) handleWeights(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		svc.getWeights(writer, req)
	case "POST":
		svc.addWeight(writer, req)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed: only GET and POST"))
	}
}

func (svc *WeightService) addWeight(writer http.ResponseWriter, req *http.Request) {

	bodyBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	newWeight := Weight{Date: time.Now()}

	err = json.Unmarshal(bodyBytes, &newWeight)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	svc.db.Create(&newWeight)

	writer.WriteHeader(http.StatusNoContent)
}

func (svc *WeightService) getWeights(writer http.ResponseWriter, req *http.Request) {

	var weights []Weight
	svc.db.Order("date").Find(&weights)

	jsonBytes, err := json.Marshal(weights)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonBytes)
}

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
