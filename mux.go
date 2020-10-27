package main

import (
	"fmt"
	"letsgo/middleware"
	"letsgo/utils"
	"letsgo/utils/router"
	"net/http"
	"time"
)


var myServeMux MyServeMux

type MyServeMux struct {
	*http.ServeMux
}

func (s *MyServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("02/Jan/2006 15:04:05")
	fmt.Printf("[%s] %s %s\n", now, r.Method, r.URL.Path)
	middleware.Middlewares.ProcessRequest(w, r)
	s.ServeMux.ServeHTTP(w, r)
	middleware.Middlewares.ProcessResponse(w, r)
}

func InitMyServeMux() {
	myServeMux.ServeMux = http.NewServeMux()
}

func GetMyServeMux() *MyServeMux {
	return &myServeMux
}

func LoadRouter(rs []router.Route) {
	router.Load(myServeMux.ServeMux, rs)
}

func ListenAndServe(addr string) {
	now := time.Now().Format("Jan 02, 2006 - 15:04:05")
	fmt.Println(now)
	fmt.Printf("Starting server at http://%s \n", addr)
	fmt.Printf("Quit the server with CTRL-BREAK.\n\n")
	http.ListenAndServe(addr, &myServeMux)
}

func StopServe() {
	err := utils.CloseDB()
	if err != nil {
		fmt.Println(err)
	}
}