package router

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler http.Handler
}

func Load(mux *http.ServeMux, rs []Route) {
	for _, route := range rs {
		mux.Handle(route.Path, route.Handler)
	}
}
