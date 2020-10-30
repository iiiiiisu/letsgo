package main

import (
	_ "github.com/go-sql-driver/mysql"
	"letsgo/cmd/demo/handler"
	"letsgo/cmd/demo/middleware"
	"letsgo/pkg"
)

func main() {
	defer pkg.StopServe()
	pkg.RegisterMiddlewares(&middleware.Middlewares)
	pkg.LoadRouter(handler.UrlPattern)
	pkg.ListenAndServe("localhost:8000")
}
