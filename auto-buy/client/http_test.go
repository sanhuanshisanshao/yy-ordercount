package client

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestHttpPost(t *testing.T) {
	url := "http://www.uuplush.com/user/refyue"
	referer := "http://www.uuplush.com/v2/index?gcid=11"
	cookie := ""

	resp, err := HttpPost(url, "", cookie, referer)
	if err != nil {
		t.Fatalf("%v", err)
	}
	f, err := strconv.ParseFloat(string(resp[1:len(resp)-1]), 32)
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Printf("%.2f", f)
}

func TestRandom(t *testing.T) {
	for i := 0; i <= 100; i++ {
		fmt.Println(rand.Intn(999-100) + 100)
	}
}

func TestHttpGet(t *testing.T) {
	url := "http://www.uuplush.com/buyorder?gcid=17&gpid=46&fieldnum=180919-39"
	cookie := "userid=4904; username=4904; ASP.NET_SessionId=5ynj4k45lrzrj21ygsxurjjo"
	resp, err := HttpGet(url, cookie)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(resp)
}
