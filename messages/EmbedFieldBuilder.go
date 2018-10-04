package messages

import "github.com/bwmarrin/discordgo"

type EmbedFieldBuilder struct {
	field *discordgo.MessageEmbedField
}

func (b *EmbedFieldBuilder) SetName(name string)  {
	if b.field == nil {
		b.field = &discordgo.MessageEmbedField{}
	}
	b.field.Name = name
}

func (b *EmbedFieldBuilder) SetValue(value string)  {
	if b.field == nil {
		b.field = &discordgo.MessageEmbedField{}
	}
	b.field.Value = value
}

func (b *EmbedFieldBuilder) SetInline(inline bool)  {
	if b.field == nil {
		b.field = &discordgo.MessageEmbedField{}
	}
	b.field.Inline = inline
}

func (b EmbedFieldBuilder) Build() *discordgo.MessageEmbedField {
	return b.field
}