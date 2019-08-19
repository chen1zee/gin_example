package gredis

import "github.com/gomodule/redigo/redis"

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		// TODO https://book.eddycjy.com/golang/gin/application-redis.html
		// 看至 四、Redis 工具包
		Dial: func() (conn redis.Conn, e error) {
			//c, err := redis.Dial("tcp", setting)
		},
		TestOnBorrow:    nil,
		MaxIdle:         0,
		MaxActive:       0,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
