package utils

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func GetUsers(s *discordgo.Session) ([]string, error) {
	var users []string

	for _, guild := range s.State.Guilds {
		for _, member := range guild.Members {
			users = append(users, member.User.ID)
		}
	}

	usersCount := len(users)
	if usersCount < 1 {
		return nil, errors.New("0 users cached")
	} else {
		return users, nil
	}
}
