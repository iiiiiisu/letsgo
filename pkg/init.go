package pkg

import (
	"fmt"
	"letsgo/pkg/conf"
)

var config conf.Config

func init() {
	InitMyServeMux()
	var err error
	config, err = conf.Read()
	if err != nil {
		panic(err)
	}
	redis := config.Redis
	dbs := config.Databases
	if len(redis) > 0 {
		InitRedis(redis[0].Network, redis[0].Addr, redis[0].Pwd)
	} else {
		panic("No Redis Config")
	}
	if len(dbs) > 0 {
		if err = InitDB("mysql", dbs[0].Username, dbs[0].Password,
			dbs[0].Host, dbs[0].Port, dbs[0].Name); err != nil {
			panic(err)
		}
	} else {
		panic("No Database Config.")
	}
	if GetRedisPool() == nil {
		fmt.Println("redis err")
	}
	if GetDB() == nil {
		fmt.Println("db err")
	}
}

func Run() {
	ListenAndServe(config.Addr)
}
