package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

func Command(s *discord.Session, m *discord.MessageCreate, url string) map[string]interface{} {
	cmd := strings.SplitAfter(m.Content, "!")
	cmd = strings.Split(cmd[1], " ")

	client := &http.Client{}

	req, err := http.NewRequest("GET", url + "api/command/" + cmd[0] + "," + cmd[1], nil)
	fmt.Println("GET", url + "api/command/" + cmd[0] + "," + cmd[1], nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("user", m.Author.Username)
	// req.Header.Add("role", )
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// var bodyString string
	
	// if resp.StatusCode == http.StatusOK {
	// 	bodyBytes, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	bodyString = string(bodyBytes)
	// 	return bodyString
	// }

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// bodyString = string(bodyBytes)
	var dat map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

	return dat
}