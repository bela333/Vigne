package debug

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/server"
)

type DebugModule struct {

}

func (DebugModule) GetName() string {
	return "debug"
}

func (DebugModule) Init(server *server.Server) error {
	cmdi, err := server.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := cmdi.(*commands.CommandsModule)
	cmd.RegisterCommand(RolesCommand{server:server})
	cmd.RegisterCommand(TestReplacingCommand{})
	return nil
}


