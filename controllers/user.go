package controllers

import (
	"log"

	discord "github.com/bwmarrin/discordgo"
)

func Admin(s *discord.Session, m *discord.MessageCreate) bool {
	perms, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		log.Println(err)
	}

	if perms&discord.PermissionAdministrator == 8 {
	    return true
	} else {
	    return false
	}
}