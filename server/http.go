package server

import "example.com/m/v2/routes"

type Http struct {
	routes.Routes
}

func (ctx *Http)HttpServer() {
	ctx.Routes.Routes()
}