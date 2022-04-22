package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	discord "github.com/bwmarrin/discordgo"
)

func Validate(s *discord.Session, m *discord.MessageCreate, u string) map[string]interface{} {
	cmd := strings.SplitAfter(m.Content, "!")
	cmd = strings.Split(cmd[1], " ")

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	

	req, err := http.NewRequest("GET", u + "api/validate/" + url.QueryEscape(strings.Join(cmd, ",")), nil)
	fmt.Println("GET", u + "api/validate/" + url.QueryEscape(strings.Join(cmd, ",")), nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("user", m.Author.Username)
	// req.Header.Add("role", )
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "This service is currently unavailable")
		return make(map[string]interface{})
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// bodyString = string(bodyBytes)
	var dat map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &dat); err != nil {
        log.Println(err)
    }
    fmt.Println(dat)

	defer resp.Body.Close()

	return dat
}