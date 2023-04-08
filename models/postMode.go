package models

import "gorm.io/gorm"

type PostModel struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}
