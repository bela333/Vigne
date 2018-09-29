package ping

type PingModule struct {

}

func (m PingModule) GetName() string {
	return "Ping"
}

func (m PingModule) Init() error {
	return nil
}


