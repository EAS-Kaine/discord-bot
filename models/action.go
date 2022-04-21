package models

import "gorm.io/gorm"

type Action struct {
	gorm.Model
	Url string
	Command string
	Permission int
	Lock bool
}

type Actions []Action