package messages

import (
	"github.com/bela333/Vigne/errors"
	"github.com/bwmarrin/discordgo"
	"time"
)

type MessageBuilder struct {
	ChannelID string
	content   string
	expiry    time.Duration

	module *MessagesModule

	embedBuilder *EmbedBuilder

	//Message was already sent
	sent bool
	//If the message was already sent, it contains the reference to the sent message
	message *discordgo.Message

	//Message shouldn't be sent
	deleted bool

	//Builder that is used to replace the current message with another one, after it gets deleted
	afterBuilder *MessageBuilder
}

func (b *MessageBuilder) getMessageSend() *discordgo.MessageSend {
	m := discordgo.MessageSend{}
	m.Content = b.content
	if b.embedBuilder != nil {
		m.Embed = b.embedBuilder.Build()
	}
	return &m
}

func (b *MessageBuilder) getMessageEdit(channelID, messageID string) *discordgo.MessageEdit {
	mSend := b.getMessageSend()
	mEdit := discordgo.MessageEdit{}

	//This probably causes a Memory Leak
	mEdit.Content = &mSend.Content
	mEdit.Embed = mSend.Embed
	mEdit.Channel = channelID
	mEdit.ID = messageID

	return &mEdit
}

func (b *MessageBuilder) afterSend(message *discordgo.Message)  {
	b.message = message
	b.sent = true
	if b.expiry != 0 {
		time.AfterFunc(b.expiry, func() {
			b.Delete()
		})
	}
}

func (b *MessageBuilder) Send() error {
	if !b.deleted {
		m := b.getMessageSend()
		msg, err := b.module.server.Session.ChannelMessageSendComplex(b.ChannelID, m)
		if err != nil {
			return err
		}
		b.afterSend(msg)
	}
	return nil
}

func (b *MessageBuilder) SetContent(content string)  {
	b.content = content
}

func (b *MessageBuilder) GetEmbedBuilder() *EmbedBuilder {
	if b.embedBuilder == nil {
		b.embedBuilder = NewEmbedBuilder()
	}
	return b.embedBuilder
}

func (b *MessageBuilder) SetExpiry(expiry time.Duration)  {
	b.expiry = expiry
}

func (b *MessageBuilder) Delete() error{
	b.deleted = true
	if b.sent {
		if b.afterBuilder != nil {
			//Remove all reactions from message
			err := b.module.server.Session.MessageReactionsRemoveAll(b.message.ChannelID, b.message.ID)
			if err != nil {
				return err
			}
			//Replace message with the next one
			err = b.ReplaceMessage(b.afterBuilder)
			if err != nil {
				return err
			}
		}else {
			err := b.module.server.Session.ChannelMessageDelete(b.message.ChannelID, b.message.ID)
			if err != nil {
				return err
			}
		}
		b.sent = false
	}
	return nil
}


//Replaces the content of the MessageBuilder with message
func (b *MessageBuilder) ReplaceMessage(message *MessageBuilder) error {
	if b.sent {
		msg, err := b.module.server.Session.ChannelMessageEditComplex(message.getMessageEdit(b.message.ChannelID, b.message.ID))
		if err != nil {
			return err
		}
		message.afterSend(msg)
	}else {
		return errors.MessageNotSent
	}
	return nil
}

//Returns a new MessageBuilder. When the current MessageBuilder gets deleted, this will replace its contents
//Only one "after MessageBuilder" can exist at a time. If one already exists, this method returns nil.
func (b *MessageBuilder) GetAfter() *MessageBuilder {
	if b.afterBuilder == nil {
		ba := MessageBuilder{}
		ba.ChannelID = b.ChannelID
		ba.module = b.module
		b.afterBuilder = &ba
		return &ba
	}else{
		return nil
	}
}