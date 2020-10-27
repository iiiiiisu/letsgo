package main

import (
	"letsgo/handler"
)

func init() {
	InitMyServeMux()
	LoadRouter(handler.UrlPattern)
}
