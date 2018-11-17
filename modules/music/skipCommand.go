package music

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/errors"
	"github.com/bela333/Vigne/messages"
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
)

type SkipCommand struct {
	server *server.Server
	module *MusicModule
}

func (SkipCommand) Check(cmd string) bool {
	return cmd == "skip"
}

func (c SkipCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	//Get db
	db, err := c.server.Database.Music()
	if err != nil {
		return err
	}
	config, err := c.server.Database.Config()
	if err != nil {
		return err
	}
	if m.ChannelID != db.GetChannel() {
		return nil
	}
	if !c.module.Player.IsPlaying {
		return errors.NotPlaying
	}
	if m.Author.ID != c.module.Player.CurrentlyPlaying.RequesterID && !config.IsMod(m.Author.ID) {
		return  errors.NotRequester
	}
	//Let's skip current song
	c.module.Player.Skip()
	return nil
}

func (SkipCommand) ShouldRemoveOriginal() bool {
	return true
}

func (SkipCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return  commands.HelpPageEntry{
		IsHidden:false,
		Usage:"",
		Description:"Skips current song.",
		Command:"skip",
	}
}
