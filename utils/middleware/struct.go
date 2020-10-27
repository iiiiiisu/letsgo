package middleware

import "net/http"

type NullMiddleWare struct {}

func (m *NullMiddleWare) ProcessRequest(w http.ResponseWriter, r *http.Request){
}

func (m *NullMiddleWare) ProcessResponse(w http.ResponseWriter, r *http.Request){

}


type MiddlewareManager struct {
	Midds []Middleware
}

func (m *MiddlewareManager) ProcessRequest(w http.ResponseWriter, r *http.Request){
	for _, midd:= range m.Midds {
		midd.ProcessRequest(w, r)
	}
}

func (m *MiddlewareManager) ProcessResponse(w http.ResponseWriter, r *http.Request){
	for _, midd:= range m.Midds {
		midd.ProcessResponse(w, r)
	}
}
