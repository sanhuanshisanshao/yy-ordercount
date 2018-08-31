package client

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type para struct {
	Gcid int `json:"gcid"`
	Gpid int `json:"gpid"`
}

func TestHttpPost(t *testing.T) {
	url := "http://www.uuplush.com/user/fieldlist"
	cookie := "safedog-flow-item=635BE5DFF566A6414BC4F418238FBA41; userid=4904; username=4904; ASP.NET_SessionId=avhcnzfz1510lv00izjp0apy"
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
