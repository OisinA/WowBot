package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
)

var beta = false

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
	file := "token.txt"
	if runtime.GOOS == "windows" {
		file = "token_beta.txt"
		beta = true
	}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(dat)
}
