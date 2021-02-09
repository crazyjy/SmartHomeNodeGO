package server

/**
 * 控制格式：
 * 		登录： LOGIN/:NAME/:TEST001
 *		控制： CONTROL/:TEST001/:OFF
 *		状态： STATUS/:TEST001/:ON
 */

import (
	"example.com/m/v2/base"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"net"
	"strings"
)

type Tcp struct {
	Base
}

/**
 * 启动 TCP 连接服务
 */
func (ctx *Base)TcpServer() {
	fmt.Println("Starting tcp server ...")
	ctx.tcpServer()
}

/**
 * TCP 服务方法
 */
func (ctx *Base)tcpServer()  {
	// 创建 listener
	listener, err := net.Listen("tcp", TCPSERVER)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return // 终止整个进程
	}
	// 建立连接池，用于广播消息
	nets := make(map[string]net.Conn)
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return // 终止程序
		}
		nets[conn.RemoteAddr().String()] = conn
		go ctx.serverStuff(conn, &nets) // 给每个连接创建一个协程处理
	}
}

/**
 * 处理客户端消息
 */
func (ctx *Base)serverStuff(conn net.Conn, nets *map[string]net.Conn) {
	for {
		buf := make([]byte, 128)
		dataLen, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				// 连接正常断开
				fmt.Println("Say goodbye !!! ")
			} else {
				fmt.Println("Error message: ", err.Error())
			}
			return // 终止该协程
		}
		data := string(buf[:dataLen])
		if strings.Contains(data, "/:") {
			info := strings.Split(data, "/:")
			if len(info) == 3 {
				fmt.Println("Call:" + data)
				ctx.checkCode(info, conn, nets)
			} else {
				fmt.Println("Error message: " + data)
			}
		} else {
			fmt.Println("Error data ??? is: " + data)
		}
	}
}

/**
 * 检查客户端消息
 */
func (ctx *Base)checkCode(info []string, conn net.Conn, nets *map[string]net.Conn) {
	code := "FAIL"
	switch info[0] {
		case "LOGIN" : // 登录注册到连接平台，建议30秒内上传一次心跳
			ctx.deviceLogin(conn.RemoteAddr().String(), info[2])
			code = "SUCCESS"
		case "CONTROL" : // 控制信息，成功返回：SUCCESS 失败返回：FAIL
			if ctx.deviceControl(nets, info[1], info[2]) {
				code = "SUCCESS"
			}
		case "STATUS" : // 状态信息
			ctx.deviceStatus(conn.RemoteAddr().String(), info[1], info[2])
			code = "SUCCESS"
	}
	_, err := conn.Write([]byte(code))
	if err != nil {
		log.Printf("Response message error: %v\n", err)
	}
}

/**
 * 处理客户端心跳方法
 */
func (ctx *Base)deviceLogin(clientId string, name string) {
	base.DeviceLink.Set(name, clientId, cache.DefaultExpiration)
}

/**
 * 处理客户端控制方法
 */
func (ctx *Base)deviceControl(nets *map[string]net.Conn, name string, code string) bool {
	data, found := base.DeviceLink.Get(name)
	if found {
		cli := (*nets)[fmt.Sprintf("%v", data)]
		_, err := cli.Write([]byte(code))
		if err != nil {
			log.Printf("Broad message failed: %v\n", err)
			return false
		}
		base.ButtonStatus.Set(name, code, cache.DefaultExpiration)
		return true
	}
	return false
}

/**
 * 处理上报状态
 */
func (ctx *Base)deviceStatus(clientId string, name string, code string) {
	base.DeviceLink.Set(name, clientId, cache.DefaultExpiration)
	base.ButtonStatus.Set(name, code, cache.DefaultExpiration)
	data, found := base.DeviceAlias.Get(name)
	if found {
		base.DeviceAlias.Set(name, data, cache.DefaultExpiration)
	}
}
