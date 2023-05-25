package container

import (
	"github.com/nsqio/go-nsq"
	"github.com/ropel12/email/config"
	"github.com/ropel12/email/pkg"
)

type Depend struct {
	Config      *config.Config
	NSQConsumer *pkg.NSQConsumer
}

func InitContainer() (*Depend, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	nsqConsumer, err := NewNSQConsumer(config)
	if err != nil {
		return nil, err
	}

	return &Depend{
		Config:      config,
		NSQConsumer: nsqConsumer,
	}, nil
}
func NewNSQConsumer(conf *config.Config) (*pkg.NSQConsumer, error) {
	nc := &pkg.NSQConsumer{}
	nc.Env = conf.NSQ
	var err error
	nsqConfig := nsq.NewConfig()
	nc.Consumer, err = nsq.NewConsumer(nc.Env.Topic, nc.Env.Channel, nsqConfig)
	if err != nil {
		return nil, err
	}

	nc.Consumer2, err = nsq.NewConsumer(nc.Env.Topic2, nc.Env.Channel2, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer3, err = nsq.NewConsumer(nc.Env.Topic3, nc.Env.Channel3, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer4, err = nsq.NewConsumer(nc.Env.Topic4, nc.Env.Channel4, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer5, err = nsq.NewConsumer(nc.Env.Topic5, nc.Env.Channel5, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer6, err = nsq.NewConsumer(nc.Env.Topic6, nc.Env.Channel6, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer7, err = nsq.NewConsumer(nc.Env.Topic7, nc.Env.Channel7, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer8, err = nsq.NewConsumer(nc.Env.Topic8, nc.Env.Channel8, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer9, err = nsq.NewConsumer(nc.Env.Topic9, nc.Env.Channel8, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer10, err = nsq.NewConsumer(nc.Env.Topic10, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer11, err = nsq.NewConsumer(nc.Env.Topic11, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer12, err = nsq.NewConsumer(nc.Env.Topic12, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	return nc, nil
}
