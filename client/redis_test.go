package client

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	addr := "0.0.0.0:6379"
	NewRedis(addr, "")

	//fmt.Println(RedisClient.Get("yy").Result())

	fmt.Println(RedisClient.HSet("5_1_1_180829_21", "total", "1000000").Result())
}
