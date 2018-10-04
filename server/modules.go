package server

import (
	"fmt"
	"github.com/bela333/Vigne/errors"
)

type Module interface {
	GetName() string
	Init(server *Server) error
}


func (s *Server) RegisterModule(m Module)  {
	fmt.Printf("Loading %s... ", m.GetName())
	m.Init(s)
	s.Modules[m.GetName()] = m
	fmt.Println("Done!")
}

func (s *Server) GetModuleByName(name string) (Module, error) {
	module, ok := s.Modules[name]
	if !ok {
		return nil, errors.NoModule
	}
	return module, nil
}