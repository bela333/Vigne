package messages

import (
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
)

type MessagesModule struct {
	server *server.Server
}

func (MessagesModule) GetName() string {
	return "messages"
}

func (m *MessagesModule) Init(server *server.Server) error {
	m.server = server
	server.Session.AddHandler(m.OnReactionAdd)
	server.Session.AddHandler(m.OnReactionRemove)
	return nil
}

func (m *MessagesModule) NewMessageCreator(channel string) *MessageCreator {
	c := MessageCreator{}
	c.ChannelID = channel
	c.module = m
	return &c
}

func (m *MessagesModule) OnReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd)  {

}

func (m *MessagesModule) OnReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove)  {

}