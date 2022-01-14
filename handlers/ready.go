package handlers

import (
	"os"

	"github.com/Raskolnikov404/goDivSource/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

var (
	prefix   string
	errorC   *color.Color
	successC *color.Color
	warningC *color.Color
)

func init() {
	prefix = os.Getenv("PREFIX")
	errorC, successC, warningC = utils.GetAllColors()
}

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	// Get guilds size
	guilds, _ := s.UserGuilds(100, "", "")
	numberOfGuilds := len(guilds)

	// Get verified state
	var verified string
	if r.User.Verified {
		verified = "✔"
	} else {
		verified = "✖"
	}

	successC.Printf("Bot conectado!\n")
	successC.Printf("Nome: %s#%s\n", r.User.Username, r.User.Discriminator)
	successC.Printf("Verificado: %s\n", verified)
	successC.Printf("Número de servidores: %d\n", numberOfGuilds)
}
