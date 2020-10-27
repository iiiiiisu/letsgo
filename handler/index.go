package handler

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"letsgo/utils/cache"
	"letsgo/utils/handler"
	"letsgo/utils/templates"
	"net/http"
)


type IndexHandler struct {
	handler.BaseHandler
	TplName string
}

func (h *IndexHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		favicon := "favicon.ico"
		http.ServeFile(w, r, favicon)
		return
	}
	data := dict{}
	cookie, _ := r.Cookie("sessionId")
	if cookie != nil {
		rConn := cache.Pool.Get()
		defer rConn.Close()
		data["username"], _ = redis.String(rConn.Do("GET", cookie.Value))
	}
	err := templates.Render(w, h.TplName, data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
