package debug

import (
	"fmt"
	"github.com/bela333/Vigne/messages"
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

type RolesCommand struct {
	server *server.Server
}

func (c RolesCommand) Check(cmd string) bool {
	return cmd == "roles"
}

func (c RolesCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	//Get configuration
	config, err := c.server.Database.Config()
	if err != nil {
		return err
	}
	if config.IsMod(m.Author.ID) {
		//User is a moderator. Let's send them the list of roles in this server
		//Get channel
		channel, err := c.server.Session.Channel(m.ChannelID)
		if err != nil {
			return err
		}
		//Get guild
		guild, err := c.server.Session.Guild(channel.GuildID)
		if err != nil {
			return err
		}
		//Create message
		message := creator.NewMessage()
		message.SetExpiry(time.Second*10)
		//Create message content
		builder := strings.Builder{}
		builder.WriteString(fmt.Sprintf("<@%s> ", m.Author.ID))
		builder.WriteString(fmt.Sprintf("Roles in %s: \n", guild.Name))
		for _, role := range guild.Roles {
			if role.Name == "@everyone" {
				continue
			}
			builder.WriteString(fmt.Sprintf("%s - %s\n", role.Name, role.ID))
		}
		message.SetContent(builder.String())

	}
	return nil
}

func (c RolesCommand) ShouldRemoveOriginal() bool {
	return true
}


