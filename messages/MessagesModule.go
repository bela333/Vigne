package messages

import (
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
)

type MessagesModule struct {
	server    *server.Server
	// Message ID -> Emoji ID -> ReactionObject
	Callbacks map[string]map[string]*ReactionObject
}

//Object that contains information about the reaction callbacks
type ReactionObject struct {
	remove func(s *server.Server, e *discordgo.MessageReactionRemove)
	add func(s *server.Server, e *discordgo.MessageReactionAdd)
}

func (r *ReactionObject) SetRemoveCallback(f func(s *server.Server, e *discordgo.MessageReactionRemove))  {
	r.remove = f
}
func (r *ReactionObject) SetAddCallback(f func(s *server.Server, e *discordgo.MessageReactionAdd))  {
	r.add = f
}

func (MessagesModule) GetName() string {
	return "messages"
}

func (m *MessagesModule) Init(s *server.Server) error {
	m.Callbacks = map[string]map[string]*ReactionObject{}
	m.server = s
	s.Session.AddHandler(m.onReactionAdd)
	s.Session.AddHandler(m.onReactionRemove)
	return nil
}

func (m *MessagesModule) NewMessageCreator(channel string) *MessageCreator {
	c := MessageCreator{}
	c.ChannelID = channel
	c.module = m
	return &c
}

func (m *MessagesModule) onReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd)  {
	//Make sure that the user isn't the current bot
	if e.UserID == s.State.User.ID {
		return
	}
	//Find callbacks for message
	handlers, ok := m.Callbacks[e.MessageID]
	if !ok {
		return
	}
	//Find callbacks for emoji
	handler, ok := handlers[e.Emoji.APIName()]
	if !ok {
		return
	}
	//Does the add callback exist?
	if handler.add != nil {
		handler.add(m.server, e)
	}
}

func (m *MessagesModule) onReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove)  {
	//Make sure that the user isn't the current bot
	if e.UserID == s.State.User.ID {
		return
	}
	//Find callbacks for message
	handlers, ok := m.Callbacks[e.MessageID]
	if !ok {
		return
	}
	//Find callbacks for emoji
	handler, ok := handlers[e.Emoji.APIName()]
	if !ok {
		return
	}
	//Does the remove callback exist?
	if handler.remove != nil {
		handler.remove(m.server, e)
	}
}