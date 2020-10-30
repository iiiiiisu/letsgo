package handler

import (
	"net/http"
)

type BaseHandler struct{}

func (h *BaseHandler) Get(w http.ResponseWriter, r *http.Request) {
}

func (h *BaseHandler) Post(w http.ResponseWriter, r *http.Request) {
}

func (h *BaseHandler) Put(w http.ResponseWriter, r *http.Request) {
}

func (h *BaseHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
