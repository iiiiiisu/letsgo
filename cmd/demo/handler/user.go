package handler

import (
	"letsgo/pkg/handler"
	"letsgo/pkg/response"
	"net/http"
)

type UserHandler struct {
	handler.BaseHandler
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	values := r.Header.Get("UrlValues")
	response.JsonResponse(w, []byte(values))
}
