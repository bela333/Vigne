package music

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/messages"
	"github.com/bwmarrin/discordgo"
)

type PlayCommand struct {

}

func (PlayCommand) Check(cmd string) bool {
	return cmd == "play"
}

func (PlayCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	return nil
}

func (PlayCommand) ShouldRemoveOriginal() bool {
	return true
}

func (PlayCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Command:"play",
		Description:"Plays music from YouTube or SoundCloud",
		Usage:"<url>",
		IsHidden: false,
	}
}
