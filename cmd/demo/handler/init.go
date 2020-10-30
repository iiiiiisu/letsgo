package handler

import (
	"letsgo/pkg/handler"
	"letsgo/pkg/router"
	"net/http"
)

type dict = map[string]interface{}

var TplPath = "cmd/demo/templates/"

var UrlPattern = []router.Route{
	router.Route{Path: "/",
		Handler: handler.NewRequestHandler(&IndexHandler{TplName: TplPath + "index.html"})},
	router.Route{Path: "/login",
		Handler: handler.NewRequestHandler(&LoginHandler{TplName: TplPath + "login.html"})},
	router.Route{Path: "/register",
		Handler: handler.NewRequestHandler(&RegisterHandler{TplName: TplPath + "register.html"})},
	router.Route{Path: "/user/(:id)",
		Handler: handler.NewRequestHandler(&UserHandler{})},
	router.Route{Path: "/static/",
		Handler: http.StripPrefix("/static/", http.FileServer(http.Dir("./cmd/demo/static")))},
}
