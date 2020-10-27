package handler

import (
	"fmt"
	"letsgo/utils/handler"
	"letsgo/utils/templates"
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
