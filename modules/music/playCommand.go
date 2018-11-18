package music

import (
	"fmt"
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/errors"
	"github.com/bela333/Vigne/messages"
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

type PlayCommand struct {
	server *server.Server
	module *MusicModule
}

func (PlayCommand) Check(cmd string) bool {
	return cmd == "play"
}

func (p PlayCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	d, err := p.server.Database.Music()
	if err != nil {
		return errors.NoMusic
	}
	if m.ChannelID != d.GetChannel() {
		m := creator.NewMessage()
		m.SetExpiry(time.Second*5)
		m.SetContent(fmt.Sprintf("Wrong channel. You should be using <#%s>", d.GetChannel()))
		return nil
	}
	message := creator.NewMessage()
	message.SetContent("Loading...")
	go func() {
		//Replace Loading... with new message after loading is done
		newMessage := message.GetAfter()
		defer message.Delete()
		//Actually add music to queue
		info, err := p.module.Player.AddMusic(strings.Join(args, " "), m.Author)
		if err == errors.InvalidExtractor {
			newMessage.SetContent("Sorry, I can't play music from this site")
			return
		}
		if err != nil {
			e, ok := err.(*errors.PublicError)
			newMessage.SetContent("Couldn't play music")
			if ok {
				newMessage.SetContent(e.PublicPart)
			}
			newMessage.SetExpiry(time.Second*10)
			return
		}
		embed := newMessage.GetEmbedBuilder()
		EmbedGenerator(embed, info, m.Author, "Added to queue", 0x26de81)

	}()
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
