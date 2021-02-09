package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

/**
 * 登录
 */
func (ctx *LoginController)Login(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	userName := json["name"]
	password := json["password"]
	if (userName == USERMAME) && (password == USERPASSWD) {
		c.JSON(http.StatusOK, ctx.GoSuccess("", ""))
	} else {
		c.JSON(http.StatusOK, ctx.GoFail("账号或密码错误", ""))
	}
}

/**
 * 登出
 */
func (ctx *LoginController)Logout(c *gin.Context) {
	/*message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})*/
}

/**
 * 获取系统时间 (10:06:50)
 */
func (ctx *LoginController)Time(c *gin.Context) {
	nowTime := ctx.NowTime()
	c.JSON(http.StatusOK, ctx.GoSuccess(nowTime[0]+":"+nowTime[1]+":"+nowTime[2], ""))
}