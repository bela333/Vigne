package music

import "github.com/bela333/Vigne/server"

type MusicModule struct {

}

func (MusicModule) GetName() string {
	return "music"
}

func (MusicModule) Init(server *server.Server) error {
	return nil
}



