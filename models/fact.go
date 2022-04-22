package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	ID string
	Subject string
	Text string
}