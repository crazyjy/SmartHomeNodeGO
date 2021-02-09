package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SwitchController struct {
	BaseController
}

/**
 * 获取首页信息
 */
func (ctx *SwitchController)Index(c *gin.Context) {
	c.JSON(http.StatusOK, ctx.GoSuccess(ctx.ReadFile(SWITCHFILE), ctx.GetButtonList()))
}

/**
 * 更新方法
 */
func (ctx *SwitchController)Update(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	data := fmt.Sprintf("%v", json["data"])
	if ctx.WriteFile(SWITCHFILE, data) {
		c.JSON(http.StatusOK, ctx.GoSuccess(data, ""))
	} else {
		c.JSON(http.StatusOK, ctx.GoFail(data, ""))
	}
}

/**
 * 控制方法
 */
func (ctx *SwitchController)Control(c *gin.Context)  {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	str := fmt.Sprintf("%v", json["data"])
	if ctx.RequestTCP(str) {
		c.JSON(http.StatusOK, ctx.GoSuccess("", ""))
	} else {
		c.JSON(http.StatusOK, ctx.GoFail("", ""))
	}
}