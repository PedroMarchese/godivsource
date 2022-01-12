package handlers

import (
	"github.com/bwmarrin/discordgo"
)

var (
	prefix string
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix = "ir!"
	if m.Content == prefix+"divulgar" {
		//
		go div(s)
	}
}

func div(s *discordgo.Session) {

}
