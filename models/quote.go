package models

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	ID string
	Text string
	Author string
}