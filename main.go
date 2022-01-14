package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Raskolnikov404/goDivSource/handlers"
	"github.com/Raskolnikov404/goDivSource/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var (
	errorC   *color.Color
	successC *color.Color
	warningC *color.Color
)

func initMain() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error on load environment variables")
	}
	fmt.Println("VARIÁVEIS DE AMBIENTE")
	fmt.Printf("TOKEN: %s\n", os.Getenv("TOKEN"))
	fmt.Printf("PREFIXO: %s\n", os.Getenv("PREFIX"))

	time.Sleep(2 * 10000)
	// utils.Clear()
	utils.DivBar()

	errorC, successC, warningC = utils.GetAllColors()
}

func main() {
	initMain()

	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		errorC.Println("error with creating discord session")
		return
	}

	// Defining intents
	session.Identify.Intents = discordgo.IntentsGuildMembers | discordgo.IntentsGuilds | discordgo.IntentsGuildMessages
	// // Adding handlers
	session.AddHandler(handlers.MessageCreate)
	// session.AddHandler(handlers.Ready)

	// Starts discord session
	err = session.Open()
	if err != nil {
		errorC.Println(err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	successC.Println("Passador Inicializando! Pressione CTRL+C para sair.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Closes discord session
	session.Close()
}
