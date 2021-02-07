package main

import (
	"context"
	"errors"
	"github.com/futurenda/google-auth-id-token-verifier"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

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

func verifyToken(db *gorm.DB, req *http.Request) (*User, error) {

	authorization, ok := req.Header["Authorization"]
	if !ok {
		return nil, errors.New("no authorization header")
	}
	token := strings.TrimPrefix(authorization[0], "bearer ")

	tokenInfo, err := verifyIdToken(token)
	if err != nil {
		return nil, err
	}

	var user User
	result := db.Where(&User{Email: tokenInfo.Email}).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user = User{
			Email: tokenInfo.Email,
			Name:  tokenInfo.Name,
		}
		db.Create(&user)
	}

	return &user, nil
}

func WithAuth(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		user, err := verifyToken(db, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), "user", *user))

		next.ServeHTTP(w, req)
	})
}
