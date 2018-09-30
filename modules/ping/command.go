package ping

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type PingCommand struct {}

func (PingCommand) Check(command string) bool {
	return command == "ping"
}

func (c *PingCommand) Action(m *discordgo.MessageCreate, args []string) error {
	fmt.Println("Pong!")
	return nil
}
