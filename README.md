# SmartHomeNodeGO

#### 介绍

用go编写的可用于智能路由器Web端管理系统。\
集成了 HTTP服务 和 TCP 连接服务。\
可给智能家居设备提供服务。\
可理解为边缘计算。

#### 特点

1. 无需安装，直接运行，所有服务会自动关联。
2. 包含了前端，后端，连接端，定时任务，文件存储。
3. 使用拓展硬盘或U盘，告别容量拮据且读写缓慢的的flash。
4. 尽量少的使用第三方包，确保能运行在大多数MIPS架构CPU的路由器上。

#### 软件架构

1. 后端使用GO开发，前端使用VueJs开发。
2. 后端框架使用gin，前端框架使用 vue-admin + element-ui。
3. 没有使用数据库，以json文件存储设备数据，简单使用了go-cache缓存。

#### 安装教程

* 如要修改本代码，请安装 GoLang 1.14+ 环境。
* 项目使用 go mod 管理包，使用以下命令：

``` bash
# 初始化 go mod 文件
go mod init example.com/m/v2
# 安装依赖
go mod tidy

# 基础运行（默认端口 8080 ）： 
go run main.go
# 指定端口： 
go run main.go -p 8081
# 系统初始化：
go run main.go -c cc
```

#### 交叉编译

``` bash
# 路由器 cpu RT-5350：
GOOS=linux GOARCH=mipsle GOMIPS=softfloat CGO_ENABLED=0 go build main.go

# 树莓派 B+：
GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 go build main.go

# 树莓派 3B：
GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build main.go

# Linux 64位：
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

*说明：
GOOS：目标平台的操作系统（darwin、freebsd、linux、windows） 
GOARCH：目标平台的体系架构（386、amd64、arm、mipsle） 
交叉编译不支持 CGO 所以要禁用它（CGO_ENABLED=0）
```

#### 使用教程

``` bash
1.  基础运行（默认端口 8080 ）： 
    ./main 
2.  指定端口：
    ./main -p 8081
3.  初始化运行：
    ./main -c cc

* 可能需要添加权限，linux 系统下请使用：
    chmod +x main
```

#### 使用说明

1. 请确保路由本机时间准确。
2. 确保防火墙允许连接服务开启的端口。
3. Web服务默认 8080 端口，TCP 连接使用 50001 端口。
4. 如需使用 80端口 请确保路由器原本后台服务已经改到别的端口。
5. *ESP8266 MINI D1 控制代码在项目 arduino/ESP8266-MINI-D1.ino

#### 项目地址
SERVER: \
github: https://github.com/crazyjy/SmartHomeNodeGO \
gitee: https://gitee.com/yjygit/SmartHomeNodeGO

H5: \
github: https://github.com/crazyjy/SmartHomeNodeH5 \
gitee: https://gitee.com/yjygit/SmartHomeNodeH5

#### 参与贡献

1.  【飞鱼智能】- 淘宝店，欢迎选购。\
https://shop223999585.taobao.com/?spm=2013.1.1000126.d21.25751acaxy0uZy