package server

import "smartHomeNode/v1/routes"

type Http struct {
	routes.Routes
}

func (ctx *Http)HttpServer() {
	ctx.Routes.Routes()
}