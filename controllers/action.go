package controllers

import (
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
		s.ChannelMessageSend(m.ChannelID, "Unsupported! Try a command! Like this: !bot actions add my_command url_to_my_api")
		s.MessageReactionAdd(m.ChannelID, m.ID, "U+1F602")
	}
}

func createAction(s *discord.Session, m *discord.MessageCreate, cmd []string, DB *gorm.DB) {
	action := models.Action{Name: cmd[3], Url: cmd[4]}
	result := DB.Create(&action)

	if result.Error != nil {
		log.Println(result.Error)
	}
}

func listActions(s *discord.Session, m *discord.MessageCreate, DB *gorm.DB) {
	var actions models.Actions
	DB.Find(&actions)
	for _, v := range actions {
		s.ChannelMessageSend(m.ChannelID, v.Name)
	}
}

func GetAction(s *discord.Session, m *discord.MessageCreate, DB *gorm.DB) (string, string) {
	cmd := strings.SplitAfter(m.Content, "!")
	cmd = strings.Split(cmd[1], " ")
	var action models.Action
	DB.Where("name = ?", cmd[0]).First(&action)
	return action.Name, action.Url
}

func Lock() {
	
}