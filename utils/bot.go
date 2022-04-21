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
	// "github.com/eas-kaine/discord-bot/models"
	"github.com/joho/godotenv"
)

const commands = `^!quiz\s|!quote\s|!image\s|!fact\s`

func Bot() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	token := os.Getenv("TOKEN")

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

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	match, _ := regexp.MatchString("(^[!])", m.Content)

	// If message is a command
	if match {
		r, err := regexp.Compile(commands)
		if err != nil {
			log.Fatal(err)
		}
		if r.MatchString(m.Content) {
			handleCommands(s, m)
		} else {
			switch m.Content {
				case "!quiz": 
					s.ChannelMessageSend(m.ChannelID, "Quiz actions: ")
				case "!quote": 
					s.ChannelMessageSend(m.ChannelID, "Quote actions: \n !quote all    !quote add")
				case "!image": 
					s.ChannelMessageSend(m.ChannelID, "Image actions: ")
				case "!fact": 
					s.ChannelMessageSend(m.ChannelID, "Fact actions: ")
				default:
					s.ChannelMessageSend(m.ChannelID, "I am not aware of this command!")
			}
		}

	// fmt.Println(m.Author.ID, m.ChannelID)

	// app, err := s.Application("966007607171121252")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(app)

	// user, err := s.GuildMember(app.GuildID, m.Author.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	//fmt.Println(s.GuildRoles())

	// fmt.Println(s.UserChannelPermissions(m.Author.ID, m.ChannelID))
	
		// default:
		// 	s.ChannelMessageSend(m.ChannelID, "I am not aware of this command!")
		// case "!createAction":
		// 	if admin(s, m) {
		// 		s.ChannelMessageSend(m.ChannelID, "Let's create an action! \n ")
				
		// 		// createAction("MOCK ACTION", 0)
		// 	} else {
		// 		s.ChannelMessageSend(m.ChannelID, "You need to be an admin to do that!")
		// 	}
		
		// case "!listActions":
		// 	if admin(s, m) {
		// 		// listActions()
		// 	} else {
		// 		s.ChannelMessageSend(m.ChannelID, "You need to be an admin to do that!")
		// 	}
		// case "!deleteAction":
		// 	if admin(s, m) {
		// 		// deleteAction()
		// 	} else {
		// 		s.ChannelMessageSend(m.ChannelID, "You need to be an admin to do that!")
		// 	}
		// }
	}
}

func handleCommands(s *discord.Session, m *discord.MessageCreate) {
	// !quiz commands
	if strings.HasPrefix(m.Content, "!quiz") && len(strings.Split(m.Content, " ")) == 2 {
		controllers.HandleQuiz(s, m)
	} else if strings.HasPrefix(m.Content, "!quote") && len(strings.Split(m.Content, " ")) == 2 {
		controllers.HandleQuote(s, m)
	} else if strings.HasPrefix(m.Content, "!image") && len(strings.Split(m.Content, " ")) == 2 {
		controllers.HandleImage(s, m)
	} else if strings.HasPrefix(m.Content, "!fact") && len(strings.Split(m.Content, " ")) == 2 {
		controllers.HandleFact(s, m)
	} else {
		s.ChannelMessageSend(m.ChannelID, "I am not aware of this command!")
	}

}