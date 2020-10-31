package handler

import (
	"letsgo/cmd/demo/models"
	"letsgo/pkg/handler"
	"net/http"
)

type LogoutHandler struct {
	handler.BaseHandler
}


func (h *LogoutHandler) Get(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("sessionId")
	if cookie != nil {
		u := models.User{}
		u.Logout(cookie.Value)
	}
	http.Redirect(w, r, "/", http.StatusFound)
	return
}