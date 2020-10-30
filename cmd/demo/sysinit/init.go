package sysinit

import (
	"letsgo/pkg"
)

func init() {
	pkg.InitDB()
	pkg.InitRedis()
	pkg.InitMyServeMux()
}
