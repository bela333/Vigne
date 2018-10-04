package welcome

import (
	"fmt"
	"github.com/bela333/Vigne/database"
	"github.com/bela333/Vigne/messages"
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"time"
)

type WelcomeModule struct {
	server *server.Server
	Database *database.Welcomer
}

func (WelcomeModule) GetName() string {
	return "welcome"
}

func (m *WelcomeModule) Init(server *server.Server) error {
	var err error
	m.Database, err = server.Database.Welcomer()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	m.server = server
	server.Session.AddHandler(m.OnJoin)
	server.Session.AddHandler(m.OnLeave)
	return nil
}

func (m *WelcomeModule) OnJoin(s *discordgo.Session, e *discordgo.GuildMemberAdd)  {
	//Send message in main channel
	creator := messages.MessageCreator{ChannelID:m.Database.GetMain()}
	before := creator.NewMessage()
	before.SetContent(fmt.Sprintf(m.Database.GetTextBefore(), e.User.ID))
	before.SetExpiry(time.Second*10)

	after := before.GetAfter()
	after.SetContent(fmt.Sprintf(m.Database.GetTextAfter(), e.User.ID))

	creator.Send(m.server)
	//Send message in "secret" channel
	if m.Database.GetSecret() != "" {
		creator := messages.MessageCreator{ChannelID:m.Database.GetSecret()}

		secret := creator.NewMessage()
		secret.SetContent(fmt.Sprintf(`%s has joined. User ID: %s
Full name: %s#%s
Mention: <@%s>.`, e.Member.User.Username, e.Member.User.ID, e.Member.User.Username, e.Member.User.Discriminator, e.Member.User.ID))

		creator.Send(m.server)
	}
}

func (m *WelcomeModule) OnLeave(s *discordgo.Session, e *discordgo.GuildMemberRemove)  {
	//Send message in "secret" channel
	if m.Database.GetSecret() != "" {
		creator := messages.MessageCreator{ChannelID:m.Database.GetSecret()}

		secret := creator.NewMessage()
		secret.SetContent(fmt.Sprintf(`%s has left. User ID: %s
Full name: %s#%s
Mention: <@%s>.`, e.Member.User.Username, e.Member.User.ID, e.Member.User.Username, e.Member.User.Discriminator, e.Member.User.ID))

		creator.Send(m.server)
	}
}