package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

func OverwatchCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	ch := make(chan string)
	usertag := strings.Replace(message, "#", "-", 1)
	go getStats(usertag, ch)
	rank := <-ch
	if rank == "" {
		s.ChannelMessageSend(m.ChannelID, "There was an error fetching your stats.")
		return
	}
	image := <-ch
	s.ChannelMessageSendEmbed(m.ChannelID, &discord.MessageEmbed{
		Title:       message,
		Description: "Rank: " + rank,
		Image: &discord.MessageEmbedImage{
			URL: image,
		},
	})
}

func getStats(user string, ch chan string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://owapi.net/api/v3/u/"+user+"/stats", nil)
	if err != nil {
		ch <- ""
		return
	}
	req.Header.Add("Accept", "text/json")
	resp, err := client.Do(req)
	if err != nil {
		ch <- ""
		return
	}
	read, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- ""
		return
	}
	var dat map[string]interface{}
	json.Unmarshal([]byte(read), &dat)
	if dat["eu"] == nil {
		ch <- ""
		return
	}
	region := dat["eu"].(map[string]interface{})
	if region == nil {
		ch <- ""
		return
	}
	stats := region["stats"].(map[string]interface{})
	if stats == nil {
		ch <- ""
		return
	}
	competitive := stats["competitive"].(map[string]interface{})
	if competitive == nil {
		ch <- ""
		return
	}
	overallStats := competitive["overall_stats"].(map[string]interface{})
	if overallStats == nil {
		ch <- ""
		return
	}
	rank := overallStats["comprank"]
	if rank == nil {
		ch <- "none"
	} else {
		ch <- fmt.Sprint(rank.(string))
	}
	ch <- overallStats["avatar"].(string)
}
