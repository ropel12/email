package pkg

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"
	"github.com/ropel12/email/config"
	"github.com/ropel12/email/entities"
	"github.com/ropel12/email/service"
	"gorm.io/gorm"
)

type NSQConsumer struct {
	Consumer   *nsq.Consumer
	Consumer2  *nsq.Consumer
	Consumer3  *nsq.Consumer
	Consumer4  *nsq.Consumer
	Consumer5  *nsq.Consumer
	Consumer6  *nsq.Consumer
	Consumer7  *nsq.Consumer
	Consumer8  *nsq.Consumer
	Consumer9  *nsq.Consumer
	Consumer10 *nsq.Consumer
	Consumer11 *nsq.Consumer
	Consumer12 *nsq.Consumer
	Consumer13 *nsq.Consumer
	Consumer14 *nsq.Consumer
	Consumer15 *nsq.Consumer
	Env        config.NSQConfig
}

func (nc *NSQConsumer) Start(sdata config.SenderConfig, conf *config.Config, db *gorm.DB) error {
	nc.Consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendEmailPendingPayment(sdata, data)
		message.Finish()
		return nil
	}))

	nc.Consumer2.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendEmailSuccessPayment(sdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer3.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendEmailCancelPayment(sdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer4.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendEmailRefundPayment(sdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer5.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		go service.SendEmailVerification(sdata, string(message.Body))
		message.Finish()
		return nil
	}))
	nc.Consumer6.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		go service.SendEmailResetPassword(sdata, string(message.Body))
		message.Finish()
		return nil
	}))
	nc.Consumer7.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		go service.SendEmailChangeEmail(sdata, string(message.Body))
		message.Finish()
		return nil
	}))
	nc.Consumer8.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendTest(sdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer9.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		go func() {
			data := entities.ReqDataQuiz{}
			json.Unmarshal(message.Body, &data)
			service.AddQuiz(data, conf.AuthQuiz)
			message.Finish()
		}()
		return nil
	}))
	nc.Consumer10.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {

		conf.AuthQuiz = string(message.Body)
		message.Finish()
		return nil
	}))
	nc.Consumer11.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendDetailCost(sdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer12.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendFinishRegister(sdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer13.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendFailRegistration(sdata, data)
		message.Finish()
		return nil
	}))
	nc.Consumer14.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.InsertSchedule(data, db)
		message.Finish()
		return nil
	}))
	nc.Consumer15.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := entities.Data{}
		json.Unmarshal(message.Body, &data)
		go service.SendMonthlyBilling(sdata, data)
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
	if err := nc.Consumer8.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer9.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer10.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer11.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer12.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer13.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer14.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
		return err
	}
	if err := nc.Consumer15.ConnectToNSQD(nc.Env.Host + ":" + nc.Env.Port); err != nil {
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
	nc.Consumer8.Stop()
	nc.Consumer9.Stop()
	nc.Consumer10.Stop()
	nc.Consumer11.Stop()
	nc.Consumer12.Stop()
	nc.Consumer13.Stop()
	nc.Consumer14.Stop()
}
