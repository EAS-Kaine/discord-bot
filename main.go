package main

import (
	"github.com/eas-kaine/discord-bot/utils"
)

func main() {


	utils.SetupDB()
	utils.Bot()

	// controllers.CreateAction(utils.DB, "MOCK ACTION", 0)
}