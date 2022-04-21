package models

import "gorm.io/gorm"

type Action struct {
	gorm.Model
	Name string
	Url string
	// Commands []string 
	// Permission int
	// Lock bool
}

type Actions []Action