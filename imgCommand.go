package main

import (
	"fmt"
	discord "github.com/bwmarrin/discordgo"
	"golang.org/x/net/html"
	"net/http"
)

func ImgCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	img := GetImage()
	fmt.Println(img)
	s.ChannelMessageSendEmbed(m.ChannelID, &discord.MessageEmbed{
		Title:       "Image",
		Description: message,
		Image: &discord.MessageEmbedImage{
			URL: GetImage(),
		},
	})
}

func GetImage() string {
	resp, _ := http.Get("https://apod.nasa.gov/apod/astropix.html")
	z := html.NewTokenizer(resp.Body)
OUTERLOOP:
	for {
		iter := z.Next()
		switch {
		case iter == html.ErrorToken:
			break OUTERLOOP
		case iter == html.StartTagToken:
			t := z.Token()

			isImage := t.Data == "img"
			if isImage {
				for _, a := range t.Attr {
					if a.Key == "src" {
						return "https://apod.nasa.gov/apod/" + a.Val
					}
				}
			}
		}
	}
	return ""
}
