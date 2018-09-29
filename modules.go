package main

type Module interface {
	GetName() string
	Init() error
}


func (s *Server) RegisterModule(m Module)  {
	m.Init()
	s.Modules = append(s.Modules, m)
}