package server

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"strings"
	"time"
)

type Cron struct {
	Base
}

/**
 * 启动服务
 */
func (ctx *Base)CronServer() {
	fmt.Println("Starting cron server ...")
	c := cron.New() 		// 创建定时任务 分钟
	spec := "* * * * *"    	// 每一分钟
	c.AddFunc(spec, func() {
		ctx.runTime()
	})
	c.Start()
	select {}
}

/**
 * 检查当前时间需要执行的任务
 */
func (ctx *Base)runTime() {
	weekDay := int(time.Now().Weekday())
	filePath := WEEKFILE + fmt.Sprintf("%v",weekDay)
	if ctx.CheckDir(filePath) {
		for _, fileList := range ctx.GetFileList(filePath) {
			path := filePath + "/" + fileList.Name()
			fileInfo := ctx.ReadFile(path)
			checkInfo := strings.Split(fileInfo,"#")
			if len(checkInfo) == 2 {
				nowTime := ctx.NowTime()
				if (nowTime[0] + ":" + nowTime[1]) == checkInfo[0] {
					arr := strings.Split(checkInfo[1],",")
					if len(arr) > 0 {
						for _, code := range arr {
							fmt.Println("Cron call: " + code)
							ctx.RequestTCP(code)
						}
					}
				}
			} else {
				fmt.Println("Cron server data error")
			}
		}
	} else {
		ctx.CreateDir(filePath)
	}
}