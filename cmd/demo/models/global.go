package models

import (
	"letsgo/pkg"
)

var db = pkg.GetDB()
var redisPool = pkg.GetRedisPool()
