package server

import (
	"github.com/bela333/Vigne/database"
	"github.com/bela333/Vigne/errors"
	"github.com/bwmarrin/discordgo"
)

type Server struct {
	Modules map[string]Module
	Session *discordgo.Session
	Database *database.Database
}

func NewServer(identifier, address, password string) (*Server, error) {
	//Create database
	var err error
	s := Server{}
	s.Modules = make(map[string]Module)
	s.Database = database.NewDatabase(identifier, address, password)
	//Get config from database
	config, err := s.Database.Config()
	//If there is no configuration ready
	if err == errors.NoConfig {
		//Create default config
		err = s.Database.CreateConfig()
		if err != nil {
			return nil, err
		}
		return nil, errors.CreatedConfig
	}
	if err != nil {
		return nil, err
	}
	s.Session, err = discordgo.New(config.Token())
	if err != nil {
		return nil, err
	}
	s.Session.ShouldReconnectOnError = true
	s.Session.StateEnabled = true
	return &s, nil
}

func (s Server) Start() error {
	err := s.Session.Open()
	if err != nil {
		return err
	}
	//TODO: Find a better method for waiting
	<- make(chan bool)
	return nil
}