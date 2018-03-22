package main

import (
	"encoding/json"
	discord "github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strings"
)

func LyricCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	split := strings.Split(message, "-")
	if len(split) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Lyrics not found.")
		return
	}
	ch := make(chan string)
	go GetLyrics(ch, split[0], split[1])
	lyrics := <-ch
	if strings.Compare(lyrics, "error") == 0 {
		s.ChannelMessageSend(m.ChannelID, "Lyrics not found.")
		return
	}
	lyricList := make([]string, len(lyrics)/1000)
	length := 0
	for {
		if length+1000 > len(lyrics) {
			lyricList = append(lyricList, lyrics[length-1000:])
			break
		}
		lyricList = append(lyricList, lyrics[length:(length+1000)])
		length += 1000
	}
	s.ChannelMessageSend(m.ChannelID, "Lyrics: ")
	for _, i := range lyricList {
		s.ChannelMessageSend(m.ChannelID, string(i))
	}
}

func GetLyrics(ch chan string, artist string, song string) {
	resp, _ := http.Get("https://api.lyrics.ovh/v1/" + artist + "/" + song)
	read, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- "error"
		return
	}
	var dat map[string]interface{}
	json.Unmarshal([]byte(read), &dat)
	_, ok := dat["error"]
	if ok {
		ch <- "error"
		return
	}
	ch <- dat["lyrics"].(string)
}
