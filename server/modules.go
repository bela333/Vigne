package server

import "github.com/bela333/Vigne/errors"

type Module interface {
	GetName() string
	Init(session *Server) error
}


func (s *Server) RegisterModule(m Module)  {
	m.Init(s)
	s.Modules[m.GetName()] = m
}

func (s *Server) GetModuleByName(name string) (Module, error) {
	module, ok := s.Modules[name]
	if !ok {
		return nil, errors.NoModule
	}
	return module, nil
}