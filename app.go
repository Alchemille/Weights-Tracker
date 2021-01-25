package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/futurenda/google-auth-id-token-verifier"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
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

func verifyIdToken(idToken string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	aud := "437796282386-o9uc6s79r134b5dkb2544ttce02piq4s.apps.googleusercontent.com"
	err := v.VerifyIDToken(idToken, []string{
		aud,
	})
	if err != nil {
		return nil, errors.New("Invalid token verified")
	}
	claimSet, err := googleAuthIDTokenVerifier.Decode(idToken)
	if err != nil {
		return nil, errors.New("Invalid token decoded")
	}

	return claimSet, nil
}

func verifyToken(writer http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		bodyBytes, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		mapToken := make(map[string]string)

		err = json.Unmarshal(bodyBytes, &mapToken)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
			return
		}

		tokenInfo, err := verifyIdToken(mapToken["id_token"])

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Bienvenue " + tokenInfo.Email))
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleRoutes(db *gorm.DB, mux *http.ServeMux) {

	svc := WeightService{db}
	mux.HandleFunc("/weights", svc.handleWeights)
	mux.HandleFunc("/", defaultRouteHdl)
	mux.HandleFunc("/verify_token", verifyToken)

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
