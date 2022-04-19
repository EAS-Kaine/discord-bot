package utils

import (
	"log"
	"os"

	discord "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Bot() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	token := os.Getenv("TOKEN")

	// Create a new Discordgo session
	disc, err := discord.New(token)
	if err != nil {
		log.Println(err)
		return
	}
	
	// Create an new Application
	ap := &discord.Application{}
	ap.Name = "TestApp"
	ap.Description = "TestDesc"
	ap, err = disc.ApplicationCreate(ap)
	log.Printf("ApplicationCreate: err: %+v, app: %+v\n", err, ap)

	// Get a specific Application by it's ID
	ap, err = disc.Application(ap.ID)
	log.Printf("Application: err: %+v, app: %+v\n", err, ap)

	// Update an existing Application with new values
	ap.Description = "Whooooa"
	ap, err = disc.ApplicationUpdate(ap.ID, ap)
	log.Printf("ApplicationUpdate: err: %+v, app: %+v\n", err, ap)

	// Create a new bot account for this application
	bot, err := disc.ApplicationBotCreate(ap.ID)
	log.Printf("BotCreate: err: %+v, bot: %+v\n", err, bot)

	// Get a list of all applications for the authenticated user
	apps, err := disc.Applications()
	log.Printf("Applications: err: %+v, apps : %+v\n", err, apps)
	for k, v := range apps {
		log.Printf("Applications: %d : %+v\n", k, v)
	}

	// Delete the application
	err = disc.ApplicationDelete(ap.ID)
	log.Printf("Delete: err: %+v\n", err)
}