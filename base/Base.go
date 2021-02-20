package base

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"smartHomeNode/v1/config"
	"time"
)

type Base struct {

}

const (
	BASEPATH = config.BASEPATH
	TCPSERVER = config.TCPSERVER
	TIMEFILE = config.TIMEFILE
	SWITCHFILE = config.SWITCHFILE
)

var (
	DeviceLink *cache.Cache
	DeviceAlias *cache.Cache
	ButtonStatus *cache.Cache
)

/**
 * 系统启动初始化
 */
func (ctx *Base)Init() {
	// 初始化应用目录
	ctx.InitDir(BASEPATH)
	ctx.InitFile(TIMEFILE)
	ctx.InitFile(SWITCHFILE)
	// 设备连接信息 3分钟失效 10分钟*清理
	DeviceLink = cache.New(3*time.Minute, 10*time.Minute)
	DeviceAlias = cache.New(3*time.Minute, 10*time.Minute)
	// 开关状态信息 90秒失效 100秒清理
	ButtonStatus = cache.New(90*time.Second, 100*time.Second)
}

/**
 * 成功返回
 */
func (ctx *Base)GoSuccess(data interface{}, info interface{}) gin.H {
	return gin.H {
		"code": "SUCCESS",
		"msg": "操作成功",
		"data": data,
		"info": info,
	}
}

/**
 * 失败返回模板
 */
func (ctx *Base)GoFail(data interface{}, info interface{}) gin.H {
	return gin.H {
		"code": "FAIL",
		"msg": "操作失败",
		"data": data,
		"info": info,
	}
}

/**
 * 初始化存储文件
 */
func (ctx *Base)InitFile(path string)  {
	if !ctx.CheckFile(path) {
		ctx.CreateFile(path)
		ctx.WriteFile(path, "[]")
	}
}

/**
 * 检查文件夹是否存在
 */
func (ctx *Base)CheckDir(path string) bool {
	s,err:=os.Stat(path)
	if err!=nil{
		return false
	}
	return s.IsDir()
}

/**
 * 检查文件是否存在
 */
func (ctx *Base)CheckFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**
 * 新建文件夹
 */
func (ctx *Base)CreateDir(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("create dir err", err)
		return false
	}
	return true
}

/**
 * 新建文件
 */
func (ctx *Base)CreateFile(path string) bool {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Println("create file err", err)
		return false
	}
	return true
}

/**
 * 读取文件内容
 */
func (ctx *Base)ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(info)
}

/**
 * 写入文件内容（快速）
 */
func (ctx *Base)WriteFile(path string, info string) bool  {
	err := ioutil.WriteFile(path, []byte(info), 0666)
	if err != nil {
		fmt.Println("write file err", err)
		return false
	}
	return true
}

/**
 * 删除文件
 */
func (ctx *Base)DelFile(path string) bool {
	if ctx.CheckFile(path) {
		err := os.Remove(path)
		if err != nil {
			fmt.Println("delete file err", err)
			return false
		}
		return true
	}
	return false
}

/**
 * TCP连接并发送信息
 */
func (ctx *Base)RequestTCP(data string) bool {
	//使用Dial建立连接
	conn, err := net.Dial("tcp", TCPSERVER)
	if err != nil {
		fmt.Println("error dialing", err.Error())
		return false
	}
	defer conn.Close()
	_, err = io.WriteString(conn, data)
	if err != nil {
		fmt.Println("write string failed", err)
		return false
	}
	// 获取 TCP server 返回值
	buf := make([]byte, 32)
	_, err = conn.Read(buf)
	if err != nil {
		return false
	}
	if string(buf[0:7]) == "SUCCESS" {
		return true
	}
	return false
}

/**
 * 获取文件夹下文件列表
 */
func (ctx *Base)GetFileList(path string) []os.FileInfo {
	files, _ := ioutil.ReadDir(path)
	return files
}

/**
 * 获取在线设备（90秒内有心跳的设备）
 */
func (ctx *Base)GetButtonList() string {
	deviceString := ""
	dataList := DeviceLink.Items()
	if len(dataList) > 0 {
		for deviceName, _:= range dataList {
			deviceString += deviceName + ","
		}
	}
	if deviceString != "" {
		return deviceString[0:len(deviceString)-1]
	}
	return deviceString
}

/**
 * 获取开关状态
 */
func (ctx *Base)GetButtonStatus() string {
	deviceString := ""
	dataList := ButtonStatus.Items()
	if len(dataList) > 0 {
		for deviceName, _:= range dataList {
			data, found := ButtonStatus.Get(deviceName)
			if found {
				deviceString += deviceName + ":" + fmt.Sprintf("%v", data) + ","
			}
		}
	}
	if deviceString != "" {
		return deviceString[0:len(deviceString)-1]
	}
	return deviceString
}

/**
 * 获取开关别名
 */
func (ctx *Base)GetButtonAlias() string {
	deviceString := ""
	dataList := DeviceAlias.Items()
	if len(dataList) > 0 {
		for deviceName, _:= range dataList {
			data, found := DeviceAlias.Get(deviceName)
			if found {
				deviceString += deviceName + ":" + fmt.Sprintf("%v", data) + ","
			}
		}
	}
	if deviceString != "" {
		return deviceString[0:len(deviceString)-1]
	}
	return deviceString
}


/**
 * 初始化文件夹
 */
func (ctx *Base)InitDir(path string) {
	if !ctx.CheckDir(path) {
		ctx.CreateDir(path)
	}
}

/**
 * 获取当前时间 ["23","05","40"]
 */
func (ctx *Base)NowTime() [3]string {
	date := [3]string{"00","00","00"}
	cstZone := time.FixedZone("CST", 8*3600)  // 东八
	date[0] = fmt.Sprintf("%v", time.Now().In(cstZone).Hour())
	date[1] = fmt.Sprintf("%v", time.Now().In(cstZone).Minute())
	date[2] = fmt.Sprintf("%v", time.Now().In(cstZone).Second())
	if len(date[0]) == 1 {
		date[0] = "0"+date[0]
	}
	if len(date[1]) == 1 {
		date[1] = "0"+date[1]
	}
	if len(date[2]) == 1 {
		date[2] = "0"+date[2]
	}
	return date
}

/**
 * 删除文件夹
 */
func (ctx *Base)DelDir(path string) bool {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("delete file err", err)
		return false
	}
	return true
}

/**
 * 获取 Http 服务端口
 */
func (ctx *Base)Args() string {
	// 定义几个变量，用于接收命令行的参数值
	var port  string
	var clean string
	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.StringVar(&port, "p", "8080", "server port")
	flag.StringVar(&clean, "c", "reserve", "init server file [-c clear]")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	if clean != "reserve" {
		ctx.DelDir(BASEPATH)
		ctx.CreateDir(BASEPATH)
	}
	return ":"+port
}
