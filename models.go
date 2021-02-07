package main

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email   string `json:"email"`
	Name    string `json:"name"`
	Weights []Weight
}

type Weight struct {
	gorm.Model
	Value  int       `json:"value"`
	Date   time.Time `json:"date"`
	UserID uint      `json:"userid"`
	User   User      `json:"user"`
}
