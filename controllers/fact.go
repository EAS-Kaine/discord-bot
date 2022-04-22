package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	discord "github.com/bwmarrin/discordgo"
	"github.com/eas-kaine/discord-bot/models"
)

func HandleFact(s *discord.Session, m *discord.MessageCreate) {
	cmd := strings.Fields(m.Content)
	// !Fact
	switch cmd[1] {
	case "random":
		randomFact(s)
	case "add":
		addFact(s, m)
	case "delete":
		deleteFact(s)
	case "all":
		allFact(s, m)
	}
}

func randomFact(s *discord.Session) {

}

func addFact(s *discord.Session, m *discord.MessageCreate) {
	if Admin(s, m) {
		values := models.Fact{Subject: "Animals.", Text: "The unicorn is the national animal of Scotland."}

		data, err := json.Marshal(values)
		if err != nil {
			log.Fatal(err)
		}
	
		response, err := http.Post(
			"http://localhost:3000/facts",
			"application/json",
			bytes.NewBuffer(data),
		)
		if err != nil {
			log.Fatal(err)
		}
	
		q := models.Fact{}
	
		json.NewDecoder(response.Body).Decode(&q)
	
		s.ChannelMessageSend(m.ChannelID, string("Sent fact!"))
	} else {
		s.ChannelMessageSend(m.ChannelID, string("You need to be an admin to do that!"))
	}
}

func deleteFact(s *discord.Session) {
	
}

func allFact(s *discord.Session, m *discord.MessageCreate) {
	response, err := http.Get("http://localhost:3000/facts")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	q := []models.Fact{}
	jsonErr := json.Unmarshal(data, &q)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for _, v := range q {
		s.ChannelMessageSend(m.ChannelID, string(v.Subject))
		s.ChannelMessageSend(m.ChannelID, string(v.Text))
	}
}