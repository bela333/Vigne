package messages

import "github.com/bela333/Vigne/server"

type MessageCreator struct {
	ChannelID string
	Builders []*MessageBuilder
}

func (c *MessageCreator) NewMessage() *MessageBuilder {
	b := &MessageBuilder{}
	b.ChannelID = c.ChannelID
	c.Builders = append(c.Builders, b)
	return b
}

func (c *MessageCreator) Send(s *server.Server) error {
	var err error
	for _, builder := range c.Builders {
		err = builder.Send(s)
		if err != nil {
			return err
		}
	}
	return nil
}