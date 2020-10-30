package handler

import (
	"fmt"
	"letsgo/cmd/demo/models"
	"letsgo/pkg/handler"
	"letsgo/pkg/response"
	"letsgo/pkg/templates"
	"net/http"
)

type RegisterHandler struct {
	handler.BaseHandler
	TplName string
}

func (h *RegisterHandler) Get(w http.ResponseWriter, r *http.Request) {
	err := templates.Render(w, h.TplName, make(map[string]interface{}))
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
	response.HtmlResponse(w, h.TplName, data)
	return
}
