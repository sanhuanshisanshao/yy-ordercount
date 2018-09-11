package client

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

type para struct {
	Gcid int `json:"gcid"`
	Gpid int `json:"gpid"`
}

func TestHttpPost(t *testing.T) {
	url := "http://www.uuplush.com/user/fieldlist"
	cookie := "safedog-flow-item=635BE5DFF566A6414BC4F418238FBA41; ASP.NET_SessionId=ymuyaerb0pgszvvki43333mf; userid=4904; username=4904"
	para := para{}
	para.Gcid = 3
	para.Gpid = 1
	b, _ := json.Marshal(&para)

	resp, err := HttpPost(url, string(b), cookie)
	fmt.Println("resp:", string(resp))

	r := make([]map[string]string, 0)

	resp = resp[1 : len(resp)-1] //去除双引号 ""
	respStr := string(resp)
	respStr = strings.Replace(respStr, `\`, "", -1) //去除转译符 \

	err = json.Unmarshal([]byte(respStr), &r)
	if err != nil {
		fmt.Println("json error: %v", err)
	}

	for _, v := range r {
		fmt.Println(v)
	}
}

type buy struct {
	OrderNum int    `json:"ordernum"`
	Gcid     int    `json:"gcid"`
	Gpid     int    `json:"gpid"`
	FieldNum string `json:"fieldnum"`
	Price    int    `json:"buyprice"`
}

func TestBuyOrder(t *testing.T) {
	URL := "http://www.uuplush.com/user/buyorder"
	cookie := ""

	para := buy{
		OrderNum: 49040000000000000 + int(time.Now().Unix()*1000) + 763,
		Gcid:     12,
		Gpid:     31,
		FieldNum: "180910-74",
		Price:    100,
	}
	//gcid=11&gpid=26&fieldnum=180910-50
	b, _ := json.Marshal(&para)

	resp, err := HttpPost(URL, string(b), cookie)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	fmt.Printf("resp: %v", string(resp))

}

func TestNewRedis(t *testing.T) {
	//fmt.Println(49040000000000000 + int(time.Now().Unix()*1000) + 763)
	fmt.Println(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 9, 0, 0, 0, time.Local))
}
