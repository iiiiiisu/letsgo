package pkg

import (
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

func InitRedis(network, addr, pwd string) {
	if network == "" {
		network = "tcp"
	}
	redisPool = &redis.Pool{
		MaxIdle:   3,
		MaxActive: 8,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, addr, redis.DialPassword(pwd))
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
