package main

import "github.com/bela333/Vigne/modules/ping"

func main() {
	s, err := NewServer("vigne", "localhost:6379", "")
	if err != nil {
		panic(err)
	}
	s.RegisterModule(ping.PingModule{})
	err = s.Start()
	if err != nil {
		panic(err)
	}

}