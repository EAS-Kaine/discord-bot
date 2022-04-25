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

	go disc.AddHandler(messageCreate)

	// Defining intents
	disc.Identify.Intents = discord.IntentsGuilds | discord.IntentsGuildMessages // discord.IntentsGuildMembers

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

func messageCreate(s *discord.Session, m *discord.MessageCreate) {

	// fmt.Println(s.GuildMembers(guild, m.Author.ID, 1000))

	// Ignore all messages created by the bot itself
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
	cmd := strings.SplitAfter(m.Content, "!")
	cmd = strings.Split(cmd[1], " ")

	if strings.HasPrefix(m.Content, "!bot") {
		if len(cmd) == 2 {
			s.ChannelMessageSend(m.ChannelID, "Try a command! Like this: !bot actions add my_command url_to_my_api")
		}
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


		// callback for quiz
		                                         
	}
}