package pkg

import (
	"fmt"
	"letsgo/pkg/middleware"
	"letsgo/pkg/router"
	"net/http"
	"time"
)

var myServeMux MyServeMux

type MyServeMux struct {
	*http.ServeMux
	midds *middleware.MiddlewareManager
}

func (s *MyServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.midds.ProcessRequest(w, r)
	s.ServeMux.ServeHTTP(w, r)
	s.midds.ProcessResponse(w, r)
}

func InitMyServeMux() {
	myServeMux.ServeMux = http.NewServeMux()
	myServeMux.midds = &middleware.MiddlewareManager{}
}

func GetMyServeMux() *MyServeMux {
	return &myServeMux
}

func LoadRouter(rs []router.Route) {
	router.Load(myServeMux.ServeMux, rs)
}

func RegisterMiddlewares(m *middleware.MiddlewareManager) {
	myServeMux.midds = m
}

func ListenAndServe(addr string) {
	now := time.Now().Format("Jan 02, 2006 - 15:04:05")
	fmt.Println(now)
	fmt.Printf("Starting server at http://%s \n", addr)
	fmt.Printf("Quit the server with CTRL-BREAK.\n\n")
	http.ListenAndServe(addr, &myServeMux)
}

func StopServe() {
	err := CloseDB()
	if err != nil {
		fmt.Println(err)
	}
}
