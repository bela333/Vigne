package roles

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/server"
)

type RolesModule struct {

}

func (RolesModule) GetName() string {
	return "roles"
}

func (RolesModule) Init(server *server.Server) error {
	cmdi, err := server.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := cmdi.(*commands.CommandsModule)
	err = cmd.RegisterCommand(RoleCommand{server:server})
	if err != nil {
		return err
	}
	return nil
}


