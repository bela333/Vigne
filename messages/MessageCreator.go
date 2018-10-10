package messages

type MessageCreator struct {
	ChannelID string
	Builders []*MessageBuilder
	module *MessagesModule
}

func (c *MessageCreator) NewMessage() *MessageBuilder {
	b := &MessageBuilder{}
	b.ChannelID = c.ChannelID
	b.module = c.module
	c.Builders = append(c.Builders, b)
	return b
}

func (c *MessageCreator) Send() error {
	var err error
	for _, builder := range c.Builders {
		err = builder.Send()
		if err != nil {
			return err
		}
	}
	return nil
}