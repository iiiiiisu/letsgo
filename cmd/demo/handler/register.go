package handler

import (
	"fmt"
	"letsgo/cmd/demo/models"
	"letsgo/pkg/handler"
	"letsgo/pkg/response"
	"letsgo/pkg/templates"
	"log"
	"net/http"
)

type RegisterHandler struct {
	handler.BaseHandler
	TplName string
}

func (h *RegisterHandler) Get(w http.ResponseWriter, r *http.Request) {
	tpl, err := templates.Prepare(h.TplName, TplPath)
	if err != nil {
		log.Println(err)
		return
	}
	err = response.HtmlResponse(w, tpl, make(map[string]interface{}))
	if err != nil {
		fmt.Println(err)
	}
}

func (h *RegisterHandler) Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	nickname := r.Form.Get("nickname")
	u := models.User{}
	data := dict{
		"username": username,
		"password": password,
		"nickname": nickname,
	}
	if u.Register(username, password, nickname) {
		cookie, _ := r.Cookie("sessionId")
		if cookie != nil {
			u.SetSession(cookie.Value)
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	tpl, err := templates.Prepare(h.TplName, TplPath)
	if err != nil {
		log.Println(err)
		return
	}
	response.HtmlResponse(w, tpl, data)
	return
}
