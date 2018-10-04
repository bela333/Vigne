package messages

import (
	"github.com/bwmarrin/discordgo"
)

type EmbedBuilder struct{
	embed *discordgo.MessageEmbed
	fields []*EmbedFieldBuilder
}

func NewEmbedBuilder() *EmbedBuilder {
	b := EmbedBuilder{}
	b.embed = &discordgo.MessageEmbed{}
	b.embed.Type = "rich"
	return &b
}

func (b EmbedBuilder) SetURL(url string)  {
	b.embed.URL = url
}

func (b EmbedBuilder) SetTitle(title string)  {
	b.embed.Title = title
}

func (b EmbedBuilder) SetDescription(description string)  {
	b.embed.Description = description
}

func (b EmbedBuilder) SetTimestamp(timestamp string)  {
	b.embed.Timestamp = timestamp
}

func (b EmbedBuilder) SetColor(color int)  {
	b.embed.Color = color
}

func (b EmbedBuilder) SetColorRGB(red, green, blue int)  {
	b.SetColor(blue | green << 8 | red << 16)
}

func (b EmbedBuilder) SetFooter(text, icon string)  {
	if b.embed.Footer == nil {
		b.embed.Footer = &discordgo.MessageEmbedFooter{}
	}
	b.embed.Footer.Text = text
	b.embed.Footer.IconURL = icon
}

func (b EmbedBuilder) SetImage(url string)  {
	if b.embed.Image == nil {
		b.embed.Image = &discordgo.MessageEmbedImage{}
	}
	b.embed.Image.URL = url
}

func (b EmbedBuilder) SetThumbnail(url string)  {
	if b.embed.Thumbnail == nil {
		b.embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	}
	b.embed.Thumbnail.URL = url
}

func (b EmbedBuilder) SetVideo(url string)  {
	if b.embed.Video == nil {
		b.embed.Video = &discordgo.MessageEmbedVideo{}
	}
	b.embed.Video.URL = url
}

func (b EmbedBuilder) SetProvider(name, url string)  {
	if b.embed.Provider == nil {
		b.embed.Provider = &discordgo.MessageEmbedProvider{}
	}
	b.embed.Provider.URL = url
	b.embed.Provider.Name = name
}

func (b EmbedBuilder) SetAuthor(name, url, icon string)  {
	if b.embed.Author == nil {
		b.embed.Author = &discordgo.MessageEmbedAuthor{}
	}
	b.embed.Author.Name = name
	b.embed.Author.URL = url
	b.embed.Author.IconURL = icon
}

func (b *EmbedBuilder) NewField() *EmbedFieldBuilder {
	f := EmbedFieldBuilder{}
	b.fields = append(b.fields, &f)
	return &f
}

func (b EmbedBuilder) Build() *discordgo.MessageEmbed {
	fieldsConverted := make([]*discordgo.MessageEmbedField, len(b.fields))
	for i, field := range b.fields {
		fieldsConverted[i] = field.Build()
	}
	b.embed.Fields = fieldsConverted
	return b.embed
}