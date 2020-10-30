package middleware

import "net/http"

type Middleware interface {
	ProcessRequest(http.ResponseWriter, *http.Request)
	ProcessResponse(http.ResponseWriter, *http.Request)
}
