package help

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/server"
)

type HelpModule struct {

}

func (HelpModule) GetName() string {
	return "help"
}

func (HelpModule) Init(server *server.Server) error {
	cmdi, err := server.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := cmdi.(*commands.CommandsModule)
	cmd.RegisterCommand(&HelpCommand{CommandsModule:cmd})
	return nil
}
