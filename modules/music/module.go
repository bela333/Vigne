package music

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/server"
)

type MusicModule struct {
	Player *MusicPlayer
}

func (MusicModule) GetName() string {
	return "music"
}

func (m *MusicModule) Init(server *server.Server) error {
	m.Player = &MusicPlayer{server:server}
	//Get command handler module
	cmdInterface, err := server.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := (cmdInterface).(*commands.CommandsModule)
	cmd.RegisterCommand(&PlayCommand{server:server, module:m})
	cmd.RegisterCommand(&SkipCommand{server:server, module:m})
	return nil
}



