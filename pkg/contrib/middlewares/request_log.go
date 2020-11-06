package middleware

import (
	"letsgo/pkg/middleware"
	"log"
	"net/http"
)

type RequestLogMidd struct {
	middleware.NullMiddleWare
}

func (m *RequestLogMidd) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	//now := time.Now().Format("02/Jan/2006 15:04:05")
	//log.Printf("[%s] %s %s\n", now, r.Method, r.URL.Path)
	log.Printf("%s %s\n", r.Method, r.URL.Path)
}
