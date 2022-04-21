package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Title string
	Url string
}