package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

func Command(s *discord.Session, m *discord.MessageCreate, u string) map[string]interface{} {
	cmd := strings.SplitAfter(m.Content, "!")
	cmd = strings.Split(cmd[1], " ")

	client := &http.Client{}

	req, err := http.NewRequest("GET", u + "api/command/" + url.QueryEscape(strings.Join(cmd, ",")), nil)
	fmt.Println("GET", u + "api/command/" + url.QueryEscape(strings.Join(cmd, ",")), nil)
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