package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

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

	currentUser := req.Context().Value("user").(User)
	newWeight := Weight{Date: time.Now(), UserID: currentUser.ID} // current date if none was provided in post request.
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

	// select weights based on the user contained in the context
	currentUser := req.Context().Value("user").(User)
	svc.db.Order("date").Where(&Weight{UserID: currentUser.ID}).Find(&weights) // API client expects ordered weights

	jsonBytes, err := json.Marshal(weights)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonBytes)
}
