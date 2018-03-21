package main

import (
	"encoding/json"
	discord "github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strings"
)

func YesOrNoCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	ch := make(chan string)
	go GetEvaluation(ch)
	answer := <-ch
	image := <-ch
	s.ChannelMessageSendEmbed(m.ChannelID, &discord.MessageEmbed{
		Title: strings.ToUpper(string(answer[0])) + string(answer[1]),
		Image: &discord.MessageEmbedImage{
			URL: image,
		},
	})
}

func GetEvaluation(ch chan string) {
	resp, _ := http.Get("https://yesno.wtf/api")
	read, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- ""
		return
	}
	var dat map[string]interface{}
	json.Unmarshal([]byte(read), &dat)
	ch <- dat["answer"].(string)
	ch <- dat["image"].(string)
}
