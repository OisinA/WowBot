package main

import (
	discord "github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
)

func DadJokeCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	ch := make(chan string)
	go getJoke(ch)
	joke := <-ch
	if joke == "" {
		s.ChannelMessageSend(m.ChannelID, "There was an error fetching your joke.")
		return
	}
	s.ChannelMessageSend(m.ChannelID, joke)
}

func getJoke(ch chan string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		ch <- ""
		return
	}
	req.Header.Add("Accept", "text/plain")
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
	ch <- string(read)
}
