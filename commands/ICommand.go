package commands

import "github.com/bwmarrin/discordgo"

type ICommand interface {
	Check(string) bool
	Action(*discordgo.MessageCreate,[]string) error
}
