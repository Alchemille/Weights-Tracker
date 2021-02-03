package main

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WeightService struct {
	db *gorm.DB
}

func (svc *WeightService) handleWeights(writer http.ResponseWriter, req *http.Request) {

	err := svc.verifyToken(req)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte(err.Error()))
		return
	}
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

	newWeight := Weight{Date: time.Now()} // current date if none was provided in post request.
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
	// API client expects ordered weights
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

func (svc *WeightService) verifyToken(req *http.Request) error {

	authorization, ok := req.Header["Authorization"]
	if !ok {
		return errors.New("no authorization header")
	}
	token := strings.TrimPrefix(authorization[0], "bearer ")

	tokenInfo, err := verifyIdToken(token)
	if err != nil {
		return err
	}

	var user User
	result := svc.db.Where(&User{Email: tokenInfo.Email}).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newUser := User{
			Email: tokenInfo.Email,
			Name:  tokenInfo.Name,
		}
		svc.db.Create(&newUser)
	}
	return nil
}
