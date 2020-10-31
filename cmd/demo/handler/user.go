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
	"net/url"
)

type UserHandler struct {
	handler.BaseHandler
	TplName string
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	urlValues := r.Header.Get("UrlValues")
	values, err := url.ParseQuery(urlValues)
	if err != nil {
		log.Println(err)
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
	pageUser := models.User{Username: values.Get("id")}
	if err = pageUser.Get(); err != nil {
		log.Println(err)
		// # 用户不存在
		pageUser.Username = ""
	}
	data := dict{
		"username": u.Username,
		"nickname": u.Nickname,
		"avatar": u.Avatar,
		"gender": u.Gender,
		"p_username": pageUser.Username,
		"p_nickname": pageUser.Nickname,
		"p_avatar": pageUser.Avatar,
		"p_gender": pageUser.Gender,
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
