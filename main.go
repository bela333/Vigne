package main

import (
	"fmt"
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/modules/debug"
	"github.com/bela333/Vigne/modules/help"
	"github.com/bela333/Vigne/modules/ping"
	"github.com/bela333/Vigne/modules/roles"
	"github.com/bela333/Vigne/modules/welcome"
	"github.com/bela333/Vigne/server"
)

//TODO: Message Reaction system
//TODO: Port over commands (music bot, event?)
//TODO: License
//TODO: Public errors

func main() {
	fmt.Print("Creating bot... ")
	s, err := server.NewServer("vigne", "localhost:6379", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	//System modules
	s.RegisterModule(&commands.CommandsModule{})
	//User modules
	s.RegisterModule(&ping.PingModule{})
	s.RegisterModule(&debug.DebugModule{})
	s.RegisterModule(&roles.RolesModule{})
	s.RegisterModule(&welcome.WelcomeModule{})
	s.RegisterModule(&help.HelpModule{})
	fmt.Print("Running bot... ")
	err = s.Start()
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")

}
