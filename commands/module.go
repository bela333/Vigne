package commands

import (
	"fmt"
	"github.com/bela333/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

type CommandsModule struct {
	Server *server.Server
	Regex *regexp.Regexp
	Commands []ICommand
}

func (module CommandsModule) GetName() string {
	return "Commands"
}

func (module *CommandsModule) Init(s *server.Server) error {
	//Get command regex from database
	config, err := s.Database.Config()
	if err != nil {
		return err
	}
	module.Regex, err = regexp.Compile(config.CommandRegex())
	if err != nil {
		return err
	}
	s.Session.AddHandler(module.handleCommands)
	module.Server = s
	return nil
}

func (module *CommandsModule) handleCommands(s *discordgo.Session, m *discordgo.MessageCreate)  {
	//Does command match?
	if module.Regex.MatchString(m.Content) {
		//Get command
		submatches := module.Regex.FindStringSubmatch(m.Content)
		command := submatches[1]
		//Get arguments for commands
		var args []string
		if len(submatches) > 2 {
			args = strings.Split(submatches[2], " ")
		}
		//Loop through every possible command
		for _, commandHandler := range module.Commands {
			//Check
			if commandHandler.Check(command){
				//If found execute action
				err := commandHandler.Action(m, args)
				//Handle error
				if err != nil {
					fmt.Printf("%s (%s) has failed to execute %s. Reason: %s\n", m.Author.Username, m.Author.ID, m.Content, err)
					return
				}
				break
			}
		}
	}
}

func (module *CommandsModule) RegisterCommand(command ICommand) error {
	module.Commands = append(module.Commands, command)
	return nil
}