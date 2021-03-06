package main

import (
	"fmt"
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/messages"
	"github.com/bela333/Vigne/modules/debug"
	"github.com/bela333/Vigne/modules/help"
	"github.com/bela333/Vigne/modules/music"
	"github.com/bela333/Vigne/modules/ping"
	"github.com/bela333/Vigne/modules/reactionMenu"
	"github.com/bela333/Vigne/modules/roles"
	"github.com/bela333/Vigne/modules/welcome"
	"github.com/bela333/Vigne/server"
)

func main() {
	fmt.Print("Creating bot... ")
	s, err := server.NewServer("vigne", "localhost:6379", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	//Service modules
	s.RegisterModule(&commands.CommandsModule{})
	s.RegisterModule(&messages.MessagesModule{})
	//User modules
	s.RegisterModule(&ping.PingModule{})
	s.RegisterModule(&debug.DebugModule{})
	s.RegisterModule(&roles.RolesModule{})
	s.RegisterModule(&welcome.WelcomeModule{})
	s.RegisterModule(&help.HelpModule{})
	s.RegisterModule(&reactionMenu.ReactionModule{})
	s.RegisterModule(&music.MusicModule{})
	fmt.Print("Running bot... ")
	err = s.Start()
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")

}
