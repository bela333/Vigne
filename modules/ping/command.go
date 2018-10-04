package ping

import (
	"fmt"
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/messages"
	"github.com/bwmarrin/discordgo"
	"time"
)

type PingCommand struct {}

func (c PingCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Description: "Pong!",
		Command:"ping",
	}
}

func (c PingCommand) ShouldRemoveOriginal() bool {
	return true
}

func (PingCommand) Check(command string) bool {
	return command == "ping"
}

func (c *PingCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	msg := creator.NewMessage()
	msg.SetExpiry(time.Second*10)
	msg.SetContent(fmt.Sprintf("<@%s> Pong!", m.Author.ID))
	return nil
}