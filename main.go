package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
)

var beta = false
var session *discord.Session

func main() {
	var err error
	session, err = discord.New("Bot " + ReadToken())
	if err != nil {
		log.Fatal(err)
		return
	}

	defer session.Close()

	if err = session.Open(); err != nil {
		log.Fatal(err)
		return
	}

	err = session.UpdateStatus(0, "~help")
	if err != nil {
		log.Fatal(err)
	}

	RegisterCommands()

	go StartWebServer(session)

	session.AddHandler(ParseCommands)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-sc
}

func ReadToken() string {
	file := "token/token.txt"
	if runtime.GOOS == "windows" {
		file = "token/token_beta.txt"
		beta = true
	}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(dat)
}

func SendMessage(s *discord.Session, channelID string, message string) {
	user, err := s.User("452445290365059072")
	if err != nil {
		return
	}
	s.ChannelMessageSendEmbed(channelID, &discord.MessageEmbed{
		Title: "WowBot",
		Thumbnail: &discord.MessageEmbedThumbnail{
			URL: user.AvatarURL("32"),
		},
		Description: message,
	})
}

func PageHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connected servers: ")
	for _, g := range session.State.Guilds {
		fmt.Fprintf(w, g.Name+"\n")
	}
}

func StartWebServer(s *discord.Session) {
	http.HandleFunc("/", PageHandle)
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		return
	}
}
