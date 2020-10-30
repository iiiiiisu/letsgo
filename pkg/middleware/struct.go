package middleware

import "net/http"

type NullMiddleWare struct{}

func (m *NullMiddleWare) ProcessRequest(w http.ResponseWriter, r *http.Request) {
}

func (m *NullMiddleWare) ProcessResponse(w http.ResponseWriter, r *http.Request) {

}

type MiddlewareManager struct {
	Midds []Middleware
}

func (m *MiddlewareManager) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	for i, _ := range m.Midds {
		m.Midds[i].ProcessRequest(w, r)
	}
}

func (m *MiddlewareManager) ProcessResponse(w http.ResponseWriter, r *http.Request) {
	length := len(m.Midds)
	for i := length - 1; i >= 0; i-- {
		m.Midds[i].ProcessResponse(w, r)
	}
}
