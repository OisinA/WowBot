package main

import (
	discord "github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	session, err := discord.New("Bot " + ReadToken())
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

	session.AddHandler(ParseCommands)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-sc
}

func ReadToken() string {
	dat, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(dat)
}
