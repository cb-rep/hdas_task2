package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Kind string `json:"kind"`
	Name string `json:"name"`
	Description string `json:"description"`
}