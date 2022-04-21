package main

import (
	"github.com/eas-kaine/discord-bot/api"
	"github.com/eas-kaine/discord-bot/utils"
)

func main() {
	go api.Run()


	utils.SetupDB()
	utils.Bot()

	// controllers.CreateAction(utils.DB, "MOCK ACTION", 0)
}