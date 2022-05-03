package controllers

import (
	"log"

	discord "github.com/bwmarrin/discordgo"
)

func Admin(s *discord.Session, u string, c string) bool {
	perms, err := s.UserChannelPermissions(u, c)
	if err != nil {
		log.Println(err)
	}

	if perms&discord.PermissionAdministrator == 8 {
	    return true
	} else {
	    return false
	}
}