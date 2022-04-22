package controllers

import (
	"fmt"
	"log"
	"strings"

	discord "github.com/bwmarrin/discordgo"
	"github.com/eas-kaine/discord-bot/models"
	"gorm.io/gorm"
)

func HandleActions(s *discord.Session, m *discord.MessageCreate, DB *gorm.DB) {
	// Actions
	cmd := strings.Fields(m.Content)
	switch cmd[1] {
	case "actions":
		if cmd[2] == "add" {
			createAction(s, m, cmd, DB)
		} else if cmd[2] == "list" {
			listActions(s, m, DB)
		}
	default:
		s.ChannelMessageSend(m.ChannelID, "unsupported!")
	}
}

func createAction(s *discord.Session, m *discord.MessageCreate, cmd []string, DB *gorm.DB) {
	// , Commands: []string{cmd[5]}
	action := models.Action{Name: cmd[3], Url: cmd[4]}
	result := DB.Create(&action)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println(action.ID)
	fmt.Println(result.RowsAffected)
}

func listActions(s *discord.Session, m *discord.MessageCreate, DB *gorm.DB) {
	var actions models.Actions
	// result := DB.Find(&actions)
	// fmt.Println(result)
	// if result.Error != nil {
	// 	log.Fatal(result.Error)
	// }
	DB.Find(&actions)
	fmt.Println(&actions)
	// fmt.Println(result.RowsAffected)
}

func GetAction(s *discord.Session, m *discord.MessageCreate, DB *gorm.DB) (string, string) {
	cmd := strings.SplitAfter(m.Content, "!")
	cmd = strings.Split(cmd[1], " ")
	var action models.Action
	DB.Where("name = ?", cmd[0]).First(&action)
	return action.Name, action.Url
}

// func CreateAction(db *gorm.DB, c string, p int) {

// 	// db.Create(&models.Action{Command: c, Permission: p})
// }

func Lock() {
	
}

// func listActions(db *gorm.DB) models.Actions {
// 	results := db.Find(&models.Actions{})

// 	return results
// }