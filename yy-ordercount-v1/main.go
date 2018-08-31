package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"yy-ordercount/server"
)

func main() {

	server.Run()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("stop server signal received")
	log.Warning("shutting down server")
}
