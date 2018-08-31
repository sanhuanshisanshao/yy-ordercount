package client

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

//建立一个redis 连接
func NewRedis(addr string, pwd string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println("pong: ", pong)
	if err != nil {
		return err
	}
	RedisClient = client
	return nil
}
