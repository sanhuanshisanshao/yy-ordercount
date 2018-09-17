package config

import (
	log "github.com/sirupsen/logrus"

	ini "gopkg.in/ini.v1"
)

//struct of simple-protel.conf
type Config struct {
	Cookie []string `ini:"cookie"`
	DDUrl  string   `ini:"dd_url"`
	Phone  string   `ini:"phone"`
}

//Read Server's Config Value from "path"
func ReadConfig(path string) (Config, error) {
	var config Config

	conf, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, path)
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config)
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}
