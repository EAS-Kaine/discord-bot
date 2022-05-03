package main

import (
	// "crypto/tls"
	// "net/http"

	"github.com/eas-kaine/discord-bot/utils"
)

func main() {
    // http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}


	utils.SetupDB()
	utils.Bot()

	// controllers.CreateAction(utils.DB, "MOCK ACTION", 0)
}