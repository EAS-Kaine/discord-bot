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

func HandleQuote(s *discord.Session, m *discord.MessageCreate) {
	cmd := strings.Fields(m.Content)
	// !quote
	switch cmd[1] {
	case "random":
		randomQuote(s, m)
	case "add":
		addQuote(s, m)
	case "delete":
		deleteQuote(s, m)
	case "all":
		allQuote(s, m)
	}
}

func randomQuote(s *discord.Session, m *discord.MessageCreate) {
}

func addQuote(s *discord.Session, m *discord.MessageCreate) {
	if Admin(s, m) {
		values := models.Quote{Text: "Life is what happens when you're busy making other plans.", Author: "John Lennon"}

    	data, err := json.Marshal(values)
		if err != nil {
    	    log.Fatal(err)
    	}

		response, err := http.Post(
			"http://localhost:3000/quotes",
			"application/json",
			bytes.NewBuffer(data),
		)
    	if err != nil {
    	    log.Fatal(err)
    	}

		q := models.Quote{}

		json.NewDecoder(response.Body).Decode(&q)

		s.ChannelMessageSend(m.ChannelID, string("Sent quote!"))
	} else {
		s.ChannelMessageSend(m.ChannelID, string("You need to be an admin to do that!"))
	}
}

func deleteQuote(s *discord.Session, m *discord.MessageCreate) {
	
}

func allQuote(s *discord.Session, m *discord.MessageCreate) {
	response, err := http.Get("http://localhost:3000/quotes")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	q := []models.Quote{}
	jsonErr := json.Unmarshal(data, &q)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for _, v := range q {
		s.ChannelMessageSend(m.ChannelID, string(v.Text))
		s.ChannelMessageSend(m.ChannelID, string(v.Author))
	}
}