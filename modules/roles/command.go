package roles

import (
	"fmt"
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/errors"
	"github.com/bela333/Vigne/messages"
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis"
	"time"
)

type RoleCommand struct {
	server *server.Server
}

func (c RoleCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Description: "Gives you a role of your choosing",
		Usage:"[role name]",
		Command:"role",
	}
}

func (RoleCommand) Check(cmd string) bool {
	return cmd == "role"
}

func (c RoleCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	roleProvider, err := c.server.Database.Roles()
	if err != nil {
		return err
	}
	if len(args) < 1 {
		//No role was specified
		helpMessage := creator.NewMessage()
		helpMessage.SetExpiry(time.Second*30)

		allRoles := roleProvider.GetAllRoles()

		embed := helpMessage.GetEmbedBuilder()
		embed.SetTitle("Available roles: ")
		description := ""
		for role := range allRoles {
			description += "--role " + role + "\n"
		}
		embed.SetDescription(description)
		helpMessage.SetContent("<@" + m.Author.ID + ">\n")
		return nil
	}
	//A role WAS specified
	role, err := roleProvider.GetRoleIDFromName(args[0])
	if err != nil {
		if err == redis.Nil {
			return errors.New("", fmt.Sprintf("Couldn't find role %s.", args[0]))
		}
		return err
	}
	allRoles := roleProvider.GetAllRoles()
	//Get channel
	channel, err := c.server.Session.State.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	//Get member of guild
	member, err := c.server.Session.State.Member(channel.GuildID, m.Author.ID)
	if err != nil {
		return err
	}
	//Remove all roles that isn't the destination role
	for _, removedRole := range allRoles {
		//For all roles in the distributable roles list...
		for _, checkedRole := range member.Roles {
			//...If it contains the currently processed role...
			//...And that role isn't the role we want to give in the end...
			//Remove the role
			if checkedRole == removedRole && removedRole != role {
				err = c.server.Session.GuildMemberRoleRemove(channel.GuildID, m.Author.ID, removedRole)
				if err != nil {
					return err
				}
				//One user can have one role only once
				break
			}
		}
	}
	//Gib role
	err = c.server.Session.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, role)
	if err != nil {
		return err
	}
	successMessage := creator.NewMessage()
	successMessage.SetContent(fmt.Sprintf("<@%s> You now have the *%s* role.", m.Author.ID, args[0]))
	successMessage.SetExpiry(time.Second*10)

	return nil
}

func (RoleCommand) ShouldRemoveOriginal() bool {
	return true
}
