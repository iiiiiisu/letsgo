package handler

import (
	"github.com/gomodule/redigo/redis"
	"letsgo/cmd/demo/models"
	"letsgo/pkg"
	"letsgo/pkg/handler"
	"letsgo/pkg/response"
	"letsgo/pkg/templates"
	"log"
	"net/http"
)

type IndexHandler struct {
	handler.BaseHandler
	TplName string
}

func (h *IndexHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		favicon := PathPrefix + "favicon.ico"
		http.ServeFile(w, r, favicon)
		return
	}
	u := models.User{}
	cookie, _ := r.Cookie("sessionId")
	if cookie != nil {
		var err error
		rConn := pkg.GetRedisConn()
		defer rConn.Close()
		if u.Username, err = redis.String(rConn.Do("GET", cookie.Value)); err != nil {
			log.Println(err)
		}
		err = u.Get()
		if err != nil {
			log.Println(err)
		}
	}
	data := dict{
		"username": u.Username,
		"nickname": u.Nickname,
		"avatar": u.Avatar,
		"gender": u.Gender,
	}
	tpl, err := templates.Prepare(h.TplName, TplPath)
	if err != nil {
		log.Println(err)
		return
	}
	err = response.HtmlResponse(w, tpl, data)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
