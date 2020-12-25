package main

import (
	"encoding/json"
	"fmt"
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

	contentType := req.Header.Get("content-type")
	if contentType != "application/json" {
		writer.WriteHeader(http.StatusBadGateway)
		writer.Write([]byte("Wrong content type: Expect application/json"))
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
	svc.db.Find(&weights)

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

func handleRoutes(db *gorm.DB) {

	svc := WeightService{db}
	http.HandleFunc("/weights", svc.handleWeights)
	http.HandleFunc("/", defaultRouteHdl)

}

func main() {
	dsn := "host=localhost dbname=weights"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	AutoMigration(db)

	handleRoutes(db)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
