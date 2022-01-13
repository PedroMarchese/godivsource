package handlers

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/Raskolnikov404/goDivSource/utils"
	"github.com/bwmarrin/discordgo"
)

var (
	prefix     string
	divMessage string
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix = "ir!"
	if m.Content == prefix+"divulgar" {
		file, err := os.Open("../files/message.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		message, _ := ioutil.ReadAll(file)
		divMessage = string(message)

		div(s, m)
	}
}

func div(s *discordgo.Session, m *discordgo.MessageCreate) {
	users, err := utils.GetUsers(s)
	if err != nil {
		panic(err)
	}

	for i, user := range users {
		open, closed := 0, 0

		channel, err := s.UserChannelCreate(user.ID)
		if err != nil {
			fmt.Printf("Erro ao tentar criar canal com o usuário %s#%s\n", user.Username, user.Discriminator)
		}

		_, err = s.ChannelMessageSend(channel.ID, divMessage)
		if err != nil {
			closed++
			fmt.Printf("[❌] DM Fechada - %s#%s\n", user.Username, user.Discriminator)
		} else {
			open++
			fmt.Printf("[✔] DM Aberta - %s#%s\n", user.Username, user.Discriminator)
		}

		progress := fmt.Sprintf("%f", math.Round(float64(i/len(users)))*100)
		os.Args[0] = "Irythill Passador [" + progress + "%]"
	}
}
