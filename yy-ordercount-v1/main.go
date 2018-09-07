package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"yy-ordercount/config"
	"yy-ordercount/server"
)

func main() {

	conf, err := config.ReadConfig("config.conf")
	if err != nil {
		log.Errorf("read config error %v", err)
		return
	}

	log.Infof("cookie: %v", conf.Cookie)

	server.Run(conf)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("stop server signal received")
	log.Warning("shutting down server")
}
