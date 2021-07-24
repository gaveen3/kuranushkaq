package middleware

import (
	iris "gopkg.in/kataras/iris.v6"
)

type headerMiddleware struct {
	Server    string
	PoweredBy string
}

func (s *headerMiddleware) Serve(ctx *iris.Context) {
	ctx.Header().Set("Server", s.Server)
	ctx.Header().Set("X-Powered-By", s.PoweredBy)
	ctx.Next()
}

//NewHeaderMiddleware *
func NewHeaderMiddleware(server, xPowerdBy string) iris.HandlerFunc {

	hmw := &headerMiddleware{Server: server, PoweredBy: xPowerdBy}

	return hmw.Serve
}
