package messages

import (
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"time"
)

type MessageBuilder struct {
	channelID string
	content string
	embed *discordgo.MessageEmbed
	expiry time.Duration
}

func (b *MessageBuilder) Send(s *server.Server) error {
	m := discordgo.MessageSend{}
	m.Content = b.content
	m.Embed = b.embed
	msg, err := s.Session.ChannelMessageSendComplex(b.channelID, &m)
	if err != nil {
		return err
	}
	if b.expiry != 0 {
		time.AfterFunc(b.expiry, func() {
			s.Session.ChannelMessageDelete(msg.ChannelID, msg.ID)
		})
	}
	return nil
}

func (b *MessageBuilder) SetContent(content string)  {
	b.content = content
}

func (b *MessageBuilder) SetEmbed(embed *discordgo.MessageEmbed)  {
	b.embed = embed
}

func (b *MessageBuilder) SetExpiry(expiry time.Duration)  {
	b.expiry = expiry
}