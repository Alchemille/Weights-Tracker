package main

import (
	"gorm.io/gorm"
	"time"
)

type Weight struct {
	gorm.Model
	Value int       `json:"value"`
	Date  time.Time `json:"date"`
}
