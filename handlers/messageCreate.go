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
	divMessage string
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	message, err := s.ChannelMessage(m.ChannelID, m.ID)
	if err != nil {
		errorC.Printf("Erro ao buscar a mensagem no %s no canal %s\n", m.ID, m.ChannelID)
	}

	prefix = os.Getenv("PREFIX")
	if message.Content == prefix+"divulgar" {
		file, err := os.Open("../files/message.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		dm, _ := ioutil.ReadAll(file)
		divMessage = string(dm)

		div(s, m)
	}
}

func div(s *discordgo.Session, m *discordgo.MessageCreate) {
	usersIDs, err := utils.GetUsers(s)
	if err != nil {
		panic(err)
	}

	for i, userID := range usersIDs {
		user, _ := s.User(userID)

		// open, closed := 0, 0

		// channel, err := s.UserChannelCreate(user.ID)
		// if err != nil {
		// 	errorC.Printf("Erro ao tentar criar canal com o usuário %s#%s\n", user.Username, user.Discriminator)
		// }

		// _, err = s.ChannelMessageSend(channel.ID, divMessage)
		// if err != nil {
		// 	closed++
		// 	errorC.Printf("[❌] DM Fechada - %s#%s\n", user.Username, user.Discriminator)
		// } else {
		// 	open++
		// 	successC.Printf("[✔] DM Aberta - %s#%s\n", user.Username, user.Discriminator)
		// }
		fmt.Printf("%s#%s\n", user.Username, user.Discriminator)

		progress := fmt.Sprintf("%f", math.Round(float64(i/len(usersIDs)))*100)
		os.Args[0] = "Irythill Passador [" + progress + "%]"
	}
}
