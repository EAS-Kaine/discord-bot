package controllers

import (
	"github.com/eas-kaine/discord-bot/models"
	"gorm.io/gorm"
)

func CreateAction(db *gorm.DB, c string, p int) {
	db.Create(&models.Action{Command: c, Permission: p})
}

func Lock() {
	
}

// func listActions(db *gorm.DB) models.Actions {
// 	results := db.Find(&models.Actions{})

// 	return results
// }