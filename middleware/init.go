package middleware

import (
	"letsgo/middleware/DynamicRouter"
	session "letsgo/middleware/Session"
	"letsgo/utils/middleware"
)

var Middlewares = middleware.MiddlewareManager{
	Midds: []middleware.Middleware{
		&session.SessionMidd{},
		&DynamicRouter.DynamicRouterParserMidd{},
	},
}
