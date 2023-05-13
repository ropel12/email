package pkg

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"
	"github.com/ropel12/email/config"
	"github.com/ropel12/email/entities"
	"github.com/ropel12/email/service"
)

type NSQConsumer struct {
	Consumer  *nsq.Consumer
	Consumer2 *nsq.Consumer
	Consumer3 *nsq.Consumer
	Consumer4 *nsq.Consumer
	Consumer5 *nsq.Consumer
	Consumer6 *nsq.Consumer
	Consumer7 *nsq.Consumer
	Env       config.NSQConfig
}

func (nc *NSQConsumer) Start(rdata config.SenderConfig) error {
	nc.Consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendEmailPendingPayment(rdata, data)
		message.Finish()
		return nil
	}))

	nc.Consumer2.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendEmailSuccessPayment(rdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer3.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendEmailCancelPayment(rdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer4.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendEmailRefundPayment(rdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer5.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		go service.SendEmailVerification(rdata, string(message.Body))
		message.Finish()
		return nil
	}))
	nc.Consumer6.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		go service.SendEmailResetPassword(rdata, string(message.Body))
		message.Finish()
		return nil
	}))
	nc.Consumer7.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		go service.SendEmailChangeEmail(rdata, string(message.Body))
		message.Finish()
		return nil
	}))
	if err := nc.Consumer.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}

	if err := nc.Consumer2.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer3.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer4.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer5.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer6.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer7.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}

	return nil
}

func (nc *NSQConsumer) Stop() {
	nc.Consumer.Stop()
	nc.Consumer2.Stop()
	nc.Consumer3.Stop()
	nc.Consumer4.Stop()
	nc.Consumer5.Stop()
	nc.Consumer6.Stop()
	nc.Consumer7.Stop()
}
