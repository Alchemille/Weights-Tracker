package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Weight struct {
	Value int `json:"value"`
	Date string `json:"date"`
}

type SliceWeights struct {
	weights []Weight
}

func newWeightsHandlers() *SliceWeights {
	return &SliceWeights{
		weights: []Weight{Weight{
			Value: 62,
			Date:  "Today",
		}},
	}
}

func (weights *SliceWeights) handleWeights(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		weights.getWeights(writer, req)
	case "POST":
		weights.addWeight(writer, req)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed: only GET and POST"))
	}
}

func (weights *SliceWeights) addWeight(writer http.ResponseWriter, req *http.Request) {

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

	var newWeight Weight
	err = json.Unmarshal(bodyBytes, &newWeight)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	weights.weights = append(weights.weights, newWeight)

}

func (weights *SliceWeights) getWeights(writer http.ResponseWriter, req *http.Request) {

	jsonBytes, err := json.Marshal(weights.weights)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
	}
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonBytes)
	//fmt.Fprintf(writer, "Hi there, I love %s!", req.URL.Path[1:])
}

func main() {

	wHandlers := newWeightsHandlers()

	http.HandleFunc("/", wHandlers.handleWeights)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic (err)
	}
}