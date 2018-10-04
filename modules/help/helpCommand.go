package help

import (
	"fmt"
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/messages"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

type HelpCommand struct {
	CommandsModule *commands.CommandsModule
}

func (HelpCommand) Check(cmd string) bool {
	return cmd == "help"
}

func (c *HelpCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	if len(args) > 0 {
		msg := creator.NewMessage()
		msg.SetExpiry(time.Second*30)
		//Find command
		var cmd commands.ICommand
		for _, command := range c.CommandsModule.Commands {
			if command.GetHelpPageEntry().Command == args[0] {
				cmd = command
				break
			}
		}
		if cmd == nil {
			msg.SetContent("Couldn't find this command")
			msg.SetExpiry(time.Second*10)
			return nil
		}
		entry := cmd.GetHelpPageEntry()
		msg.SetContent(fmt.Sprintf(`--%s %s
	%s`, entry.Command, entry.Usage, entry.Description))
	}else {
		msg := creator.NewMessage()
		msg.SetExpiry(time.Minute)
		textBuilder := strings.Builder{}
		textBuilder.WriteString("Available commands: \n")
		for _, command := range c.CommandsModule.Commands {
			entry := command.GetHelpPageEntry()
			textBuilder.WriteString(fmt.Sprintf("--%s %s\n", entry.Command, entry.Usage))
		}
		msg.SetContent(textBuilder.String())
	}
	return nil
}

func (HelpCommand) ShouldRemoveOriginal() bool {
	return true
}

func (HelpCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Description: "Returns all available commands",
		Usage: "[command name]",
		Command: "help",
	}
}

