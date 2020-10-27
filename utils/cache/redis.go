package cache

import (
	"github.com/gomodule/redigo/redis"
)

var Pool = &redis.Pool{
	MaxIdle: 3,
	MaxActive: 8,
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp",
			"localhost:6379", redis.DialPassword("123456"))
		if err != nil {
			return nil, err
		}
		return c, nil
	},
}