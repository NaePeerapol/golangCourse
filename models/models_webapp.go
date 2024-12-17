package models

import "gorm.io/gorm"

type Person struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

type Dogs struct {
	gorm.Model
	DogId int `json:"dog_id"`
	DogName string `json:"dog_name"`
}