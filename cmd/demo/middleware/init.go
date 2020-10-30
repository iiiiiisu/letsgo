package middleware

import (
	"letsgo/pkg/middleware"
)

var Middlewares = middleware.MiddlewareManager{
	Midds: []middleware.Middleware{
		&SessionMidd{},
		&DynamicRouterParserMidd{},
		&RequestLogMidd{},
	},
}
