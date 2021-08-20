package middleware

import (
	iris "gopkg.in/kataras/iris.v6"
)

type headerMw struct {
	Server    string
	PoweredBy string
}

func (s *headerMw) Serve(ctx *iris.Context) {
	ctx.Header().Set("Server", s.Server)
	ctx.Header().Set("X-Powered-By", s.PoweredBy)
	ctx.Next()
}

//NewServeHeader *
func NewServeHeader(server, xPowerdBy string) iris.HandlerFunc {

	hmw := &headerMw{Server: server, PoweredBy: xPowerdBy}

	return hmw.Serve
}
