package debug

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/messages"
	"github.com/bwmarrin/discordgo"
	"time"
)

type TestReplacingCommand struct {

}

func (TestReplacingCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Description: "A debugging command",
		Command: "replace",
	}
}

func (TestReplacingCommand) Check(cmd string) bool {
	return cmd == "replace"
}

func (TestReplacingCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	m1 := creator.NewMessage()
	m1.SetContent("Hello")
	m1.SetExpiry(time.Second*5)
	m2 := m1.GetAfter()
	m2.SetContent("World!")
	m2.SetExpiry(time.Second*10)
	return nil
}

func (TestReplacingCommand) ShouldRemoveOriginal() bool {
	return false
}
