package help

import (
	"fmt"
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/messages"
	"github.com/bwmarrin/discordgo"
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
			//TODO: Public error
			msg.SetContent("Couldn't find this command")
			msg.SetExpiry(time.Second*10)
			return nil
		}
		entry := cmd.GetHelpPageEntry()
		embed := msg.GetEmbedBuilder()
		embed.SetColorRGB(0,255,255)

		field := embed.NewField()
		field.SetName(fmt.Sprintf("--%s %s", entry.Command, entry.Usage))
		field.SetValue(entry.Description)
		field.SetInline(false)

	}else {
		msg := creator.NewMessage()
		msg.SetExpiry(time.Minute)

		embed := msg.GetEmbedBuilder()
		embed.SetTitle("Available commands: ")
		embed.SetColorRGB(255,255,0)
		for _, command := range c.CommandsModule.Commands {
			entry := command.GetHelpPageEntry()
			if entry.IsHidden {
				continue
			}
			field := embed.NewField()
			field.SetName(fmt.Sprintf("--%s %s\n", entry.Command, entry.Usage))
			field.SetValue(entry.Description)
			field.SetInline(false)
		}
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

