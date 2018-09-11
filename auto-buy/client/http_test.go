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
