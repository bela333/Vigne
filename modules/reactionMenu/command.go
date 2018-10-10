package reactionMenu

import (
	"fmt"
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/messages"
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"time"
)

type ReactionCommand struct {
	server *server.Server
}

func (c ReactionCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Description: "Some reaction demo.",
		Command:"react",
		IsHidden:true,
	}
}

func (c ReactionCommand) ShouldRemoveOriginal() bool {
	return true
}

func (ReactionCommand) Check(command string) bool {
	return command == "react"
}

func (c *ReactionCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	msg := creator.NewMessage()
	msg.SetContent("Click the button!")
	msg.SetExpiry(time.Second*10)

	//http://graphemica.com/
	//"UTF-8 (hex)"
	handler := msg.GetReactionHandler("\xf0\x9f\x85\xb1") //B


	handler.SetAddCallback(func(s *server.Server, e *discordgo.MessageReactionAdd) {
		fmt.Println("B")
	})
	return nil
}