package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	discord "github.com/bwmarrin/discordgo"
)

func Command(s *discord.Session, m string, user string, id string, u string, c string) map[string]interface{} {
	cmd := strings.SplitAfter(m, "!")
	cmd = strings.Split(cmd[1], " ")

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequest("GET", u + "api/command/" + url.QueryEscape(strings.Join(cmd, ",")), nil)
	if err != nil {
		log.Println(err)
	}
	// Request header
	req.Header.Add("user", user)
	req.Header.Add("role", strconv.FormatBool(Admin(s, id, c)))

	// GET request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
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

	defer resp.Body.Close()

	return dat
}