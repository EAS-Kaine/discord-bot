package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	discord "github.com/bwmarrin/discordgo"

	"github.com/eas-kaine/discord-bot/controllers"
	"github.com/joho/godotenv"
)

var guild string

// const commands = `^!bot\s|^!quiz\s|!quote\s|!image\s|!fact\s`

func Bot() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	token := os.Getenv("TOKEN")
	guild = os.Getenv("GUILD")

	// Create a new Discordgo session
	disc, err := discord.New("Bot " + token)
	if err != nil {
		log.Println(err)
		return
	}

	disc.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	disc.Identify.Intents = discord.IntentsGuilds | discord.IntentsGuildMessages
	// discord.IntentsGuilds | discord.IntentsGuildMessages | discord.IntentsGuildMembers

	// Open a websocket connection to Discord and begin listening.
	err = disc.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	
	disc.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discord.Session, m *discord.MessageCreate) {

	// fmt.Println(s.GuildMembers(guild, m.Author.ID, 1000))

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	match, _ := regexp.MatchString("(^[!])", m.Content)

	// If message is a command
	if match {
		handleCommands(s, m)
	}
}

func handleCommands(s *discord.Session, m *discord.MessageCreate) {
	if strings.HasPrefix(m.Content, "!bot") {
		controllers.HandleActions(s, m, DB)
	} else {
		name, url := controllers.GetAction(s, m, DB)
		fmt.Println(name, url)
		fmt.Println(controllers.Validate(s, m, url))
		data := controllers.Validate(s, m, url)
		if  data["status_message"] == "valid_command" {
			data := controllers.Command(s, m, url)
			msg, ok := data["discord_message"].(string)
			if !ok {
				log.Println()
			}
			s.ChannelMessageSend(m.ChannelID, msg)
		} else if msg, _ := data["discord_message"].(string); msg != "" {
			s.ChannelMessageSend(m.ChannelID, msg)
		} 
		if msg, ok := data["discord_message_complex"].(discord.MessageSend); !ok {
			fmt.Println("Couldn't get discord_message_complex?")
		} else if msg.Content != "" {
			s.ChannelMessageSendComplex(m.ChannelID, &msg)
		}

		//!bot actions list
		                                         
	}
}