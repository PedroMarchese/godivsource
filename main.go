package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("error on load environment variables")
	}
}

func main() {
	Init()

	session, err := discordgo.New("Bot " + "TOKEN")
	if err != nil {
		fmt.Println("error with creating discord session")
		return
	}

	// Defining intents
	session.Identify.Intents = discordgo.IntentsGuildMembers | discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	// Adding handlers
	// session.AddHandler()

	// Starts discord session
	err = session.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Closes discord session
	session.Close()
}
