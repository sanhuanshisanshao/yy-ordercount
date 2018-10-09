package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"yy-ordercount/auto-buy/baseinfo"
	"yy-ordercount/auto-buy/config"
	"yy-ordercount/auto-buy/user"
)

func Ping(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("OK"))
}

func Stop(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("success"))
	log.Warnf("start stop auto-buy server")
	os.Exit(0)
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	fmt.Println("set log farmat")
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	conf, err := config.ReadConfig("config.conf")
	if err != nil {
		return
	}

	log.Infof("cookie: %v", conf.Cookie)

	//数据初始化
	user.NewUsers(conf.Cookie, conf.DDUrl, conf.Phone)
	baseinfo.NewAreaIds()
	baseinfo.NewFieldInfo()

	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/stop", Stop)

	log.Fatal(http.ListenAndServe("0.0.0.0:9998", nil))
}
