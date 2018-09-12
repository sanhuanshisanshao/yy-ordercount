package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"yy-ordercount/auto-buy/baseinfo"
	"yy-ordercount/auto-buy/config"
	"yy-ordercount/auto-buy/user"
)

func Ping(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("OK"))
}

func SetCookie(resp http.ResponseWriter, req *http.Request) {
	cookie := req.FormValue("cookie")
	log.Infof("set cookie %v success", cookie)
	if len(cookie) > 0 {
		user.UniqueUsers.Add(cookie)
	}
	resp.Write([]byte("success"))
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
	user.NewUsers(conf.Cookie)
	baseinfo.NewAreaIds()
	baseinfo.NewFieldInfo()

	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/setcookie", SetCookie)

	log.Fatal(http.ListenAndServe("0.0.0.0:9998", nil))
}
