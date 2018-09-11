package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"yy-ordercount/auto-buy/baseinfo"
	"yy-ordercount/auto-buy/user"
)

func Ping(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("OK"))
}

func SetCookie(resp http.ResponseWriter, req *http.Request) {
	cookie := req.FormValue("cookie")
	logrus.Infof("set cookie %v success", cookie)
	if len(cookie) > 0 {
		user.UniqueUsers.Add(cookie)
	}
	resp.Write([]byte("success"))
}

func main() {
	//数据初始化
	user.NewUsers()
	baseinfo.NewAreaIds()
	baseinfo.NewFieldInfo()

	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/setcookie", SetCookie)

	log.Fatal(http.ListenAndServe("0.0.0.0:9998", nil))
}
