package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type TimeController struct {
	BaseController
}

/**
 * 获取首页信息
 */
func (ctx *TimeController)Index(c *gin.Context) {
	c.JSON(http.StatusOK, ctx.GoSuccess(ctx.ReadFile(TIMEFILE), ctx.ReadFile(SWITCHFILE)))
}

/**
 * 更新内容
 */
func (ctx *TimeController)Update(c *gin.Context)  {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	data := fmt.Sprintf("%v", json["data"])
	if ctx.WriteFile(TIMEFILE, data) {
		c.JSON(http.StatusOK, ctx.GoSuccess(data, ""))
	} else {
		c.JSON(http.StatusOK, ctx.GoFail(data, ""))
	}
}

/**
 * 生成定时任务文件
 */
func (ctx *TimeController)Control(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	data := fmt.Sprintf("%v", json["data"])
	door := fmt.Sprintf("%v", json["door"])
	arr := strings.Split(data, "&")
	for key, value := range arr {
		info := strings.Split(value, "@")
		if len(info) == 2 {
			weekPath := WEEKFILE + fmt.Sprintf("%v", key) + "/"
			filePath := weekPath + info[0]
			if door == "0" || info[1] == "DEL" {
				ctx.DelFile(filePath)
			}
			if door == "1" && info[1] != "DEL" {
				ctx.InitDir(weekPath)
				ctx.WriteFile(filePath, info[1])
			}
		}
	}
	c.JSON(http.StatusOK, ctx.GoSuccess("", ""))
}