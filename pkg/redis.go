package pkg

import (
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

func InitRedis() {
	redisPool = &redis.Pool{
		MaxIdle:   3,
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
}

func GetRedisPool() *redis.Pool {
	return redisPool
}

func GetRedisConn() (conn redis.Conn) {
	conn = redisPool.Get()
	return
}
