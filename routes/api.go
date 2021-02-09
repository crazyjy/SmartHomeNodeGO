package routes

import (
	"example.com/m/v2/base"
	"example.com/m/v2/config"
	"example.com/m/v2/http/Controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Login Controllers.LoginController
	Switch Controllers.SwitchController
	Time Controllers.TimeController
	Button Controllers.ButtonController
)

type Routes struct {
	base.Base
}

/**
 * 路由
 */
func (stx *Routes)Routes() {
	fmt.Println("Starting http server ...")
	gin.SetMode(config.GINMODE)
	router := gin.Default()
	// 前端页面
	router.LoadHTMLGlob("./index.html")
	router.Static("static/css", "static/css").Static("static/js", "static/js").
		Static("static/fonts", "static/fonts").Static("static/img", "static/img")
	router.GET("./", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// 登录方法
	admin := router.Group("/api/admin")
	{
		admin.POST("/login", Login.Login)
		admin.POST("/logout", Login.Logout)
		admin.GET("/time", Login.Time)
	}

	// 控制方法
	api := router.Group("/api")
	{
		/*--------- 开关控制 ---------*/
		api.GET("/switch/index", Switch.Index)
		api.POST("/switch/update", Switch.Update)
		api.POST("/switch/control", Switch.Control)

		/*--------- 定时控制 ---------*/
		api.GET("/time/index", Time.Index)
		api.POST("/time/update", Time.Update)
		api.POST("/time/control", Time.Control)

		/*--------- 按钮控制 ---------*/
		api.GET("/button/index", Button.Index)
		api.POST("/button/control", Button.Control)
		api.GET("/button/control/:name/:control", Button.UrlControl)
	}

	router.Run(stx.Args())
}