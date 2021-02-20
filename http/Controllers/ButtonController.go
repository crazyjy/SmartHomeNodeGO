package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	"smartHomeNode/v1/base"
)

type ButtonController struct {
	BaseController
}

/**
 * 获取首页信息
 */
func (ctx *ButtonController)Index(c *gin.Context) {
	c.JSON(http.StatusOK, ctx.GoSuccess(ctx.GetButtonStatus(), ctx.GetButtonAlias()))
}

/**
 * 控制方法
 */
func (ctx *ButtonController)Control(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	str := fmt.Sprintf("%v", json["data"])
	if ctx.RequestTCP(str) {
		name := json["name"]
		alias := json["alias"]
		if name != nil && alias != nil {
			base.DeviceAlias.Set(fmt.Sprintf("%v", name), fmt.Sprintf("%v", alias), cache.DefaultExpiration)
		}
		c.JSON(http.StatusOK, ctx.GoSuccess("", ""))
	} else {
		c.JSON(http.StatusOK, ctx.GoFail("", ""))
	}
}

/**
 * 浏览器控制方法 http://server.com/api/button/control/TEST001/ON
 */
func (ctx *ButtonController)UrlControl(c *gin.Context)  {
	name := c.Param("name")
	control := c.Param("control")
	if ctx.RequestTCP("CONTROL/:"+name+"/:"+control) {
		c.JSON(http.StatusOK, ctx.GoSuccess("", ""))
	} else {
		c.JSON(http.StatusOK, ctx.GoFail("", ""))
	}
}