package main

import (
	"example.com/m/v2/base"
	"example.com/m/v2/server"
)

var b base.Base		// 基础服务
var t server.Tcp	// TCP服务
var c server.Cron 	// 定时任务
var h server.Http	// HTTP服务

func main() {
	b.Init()
	go func() {
		t.TcpServer()
	}()
	go func() {
		c.CronServer()
	}()
	go func() {
		h.HttpServer()
	}()
	select {}
}