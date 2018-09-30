package main

import (
	"github.com/bela333/Vigne/commands"
	"github.com/bela333/Vigne/modules/ping"
	"github.com/bela333/Vigne/server"
)

//TODO: PingCommand system
//TODO: License

func main() {
	s, err := server.NewServer("vigne", "localhost:6379", "")
	if err != nil {
		panic(err)
	}
	//System modules
	s.RegisterModule(&commands.CommandsModule{})
	//User modules
	s.RegisterModule(&ping.PingModule{})
	err = s.Start()
	if err != nil {
		panic(err)
	}

}