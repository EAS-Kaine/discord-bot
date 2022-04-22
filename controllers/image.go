package controllers

import (
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

func HandleImage(s *discord.Session, m *discord.MessageCreate) {
	cmd := strings.Fields(m.Content)
	// !image
	switch cmd[1] {
	case "random":
		randomImage(s)
	case "add":
		addImage(s)
	case "delete":
		deleteImage(s)
	case "all":
		allImage(s)
	}
}

func randomImage(s *discord.Session) {

}

func addImage(s *discord.Session) {
	
}

func deleteImage(s *discord.Session) {
	
}

func allImage(s *discord.Session) {
	
}