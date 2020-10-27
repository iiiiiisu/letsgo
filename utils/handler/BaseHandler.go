package handler

import (
	"fmt"
	"net/http"
)


type BaseHandler struct {}

func (h *BaseHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get")
}

func (h *BaseHandler) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post")
}

func (h *BaseHandler) Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Put")
}

func (h *BaseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete")
}