package handler

import (
	"letsgo/utils/handler"
	"net/http"
)

type UserHandler struct {
	handler.BaseHandler
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	values := r.Header.Get("UrlValues")
	w.Write([]byte(values))
}