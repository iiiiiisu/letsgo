package handler

import "net/http"

type AbstractHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type ReqeustHandler struct {
	AbstractHandler
}

func (rh *ReqeustHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rh.Get(w, r)
	case "POST":
		rh.Post(w, r)
	case "PUT":
		rh.Put(w, r)
	case "DELETE":
		rh.Delete(w, r)
	}
}

func NewRequestHandler(a AbstractHandler) *ReqeustHandler {
	return &ReqeustHandler{
		AbstractHandler: a,
	}
}
