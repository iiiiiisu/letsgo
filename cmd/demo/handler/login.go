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

type LoginHandler struct {
	handler.BaseHandler
	TplName string
}

func (h *LoginHandler) Get(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"username": "",
		"password": "",
	}
	tpl, err := templates.Prepare(h.TplName, TplPath)
	if err != nil {
		log.Println(err)
		return
	}
	err = response.HtmlResponse(w, tpl, data)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *LoginHandler) Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	data := map[string]interface{}{
		"username": username,
		"password": password,
	}
	if username == "" || password == "" {
		data["err_msg"] = "请输入账号和密码"
	} else {
		u := models.User{}
		if u.Login(username, password) {
			cookie, _ := r.Cookie("sessionId")
			if cookie != nil {
				u.SetSession(cookie.Value)
			}
			http.Redirect(w, r, "/", http.StatusFound)
			return
		} else {
			data["err_msg"] = "账号或密码错误"
		}
	}
	tpl, err := templates.Prepare(h.TplName, TplPath)
	if err != nil {
		log.Println(err)
		return
	}
	err = response.HtmlResponse(w, tpl, data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
