package ping

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/server"
)

type PingModule struct {

}

func (m PingModule) GetName() string {
	return "Ping"
}

func (m *PingModule) Init(s *server.Server) error {
	//Get command handler module
	cmdInterface, err := s.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := (cmdInterface).(*commands.CommandsModule)
	cmd.RegisterCommand(&PingCommand{})
	return nil
}


