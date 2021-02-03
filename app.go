package main

import (
	"errors"
	"fmt"
	"github.com/futurenda/google-auth-id-token-verifier"
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
	handler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}).Handler(mux)
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

}
