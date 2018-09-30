package commands

import (
	"github.com/bela333/Vigne/messages"
	"github.com/bwmarrin/discordgo"
)

type ICommand interface {
	Check(string) bool
	Action(m *discordgo.MessageCreate,args []string, creator *messages.MessageCreator) error
}
