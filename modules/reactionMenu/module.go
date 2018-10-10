package reactionMenu

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/server"
)

type ReactionModule struct {

}

func (m ReactionModule) GetName() string {
	return "reactionMenu"
}

func (m *ReactionModule) Init(s *server.Server) error {
	//Get command handler module
	cmdInterface, err := s.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := (cmdInterface).(*commands.CommandsModule)
	cmd.RegisterCommand(&ReactionCommand{server:s})
	return nil
}


