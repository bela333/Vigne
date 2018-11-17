package music

import (
	"fmt"
	"github.com/bela333/Vigne/messages"
	"github.com/bwmarrin/discordgo"
)

func EmbedGenerator(embed *messages.EmbedBuilder, info *Music, author *discordgo.User, title string, color int)  {
	embed.SetTitle("**" + title + "**")
	nameField := embed.NewField()
	nameField.SetName("Title")
	nameField.SetValue(info.Title)
	embed.SetURL(info.URL)
	if info.Uploader != "" {
		uploaderField := embed.NewField()
		uploaderField.SetName("By")
		uploaderField.SetValue(info.Uploader)
	}
	timeField := embed.NewField()
	timeField.SetName("Length")
	timeField.SetValue(FormatTime(info.Duration))
	embed.SetColor(color)
	embed.SetImage(info.Thumbnail)
	embed.SetAuthor(fmt.Sprintf("Requested by %s#%s", author.Username, author.Discriminator),
		"",
		fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", author.ID, author.Avatar))
}