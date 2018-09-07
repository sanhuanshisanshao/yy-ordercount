package config

import (
	log "github.com/Sirupsen/logrus"

	ini "gopkg.in/ini.v1"
)

//struct of simple-protel.conf
type Config struct {
	HttpPort  string `ini:"http_port"`
	RedisAddr string `ini:"redis_addr"`
	RedisPwd  string `ini:"redis_pwd"`
	Cookie    string `ini:"cookie"`

	//LOG
	LogLevel    string `ini:"log_level"`
	LogDirWin   string `ini:"log_dir_win"`
	LogDirLinux string `ini:"log_dir_linux"`
	LogPrefix   string `ini:"log_prefix"`
}

//Read Server's Config Value from "path"
func ReadConfig(path string) (Config, error) {
	var config Config

	conf, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, path)
	//conf, err := ini.Load(path)
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
