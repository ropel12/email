package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/ropel12/email/config/container"
)

func main() {
	con, err := container.InitContainer()
	if err != nil {
		log.Fatalf("[FATAL] Failed to inject dependency: %v", err)
		return
	}

	go func() {
		log.Println("[INFO] Starting Service Consumer")
		if err := con.NSQConsumer.Start(con.Config.Sender, con.Config); err != nil {
			log.Fatalf("[FATAL] Failed to start NSQ Consumer: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	con.NSQConsumer.Stop()
	log.Println("[INFO]  Email Service Consumer Stopped")
}
