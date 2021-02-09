package config

/**
 * 配置项
 */
const (
	USERMAME = "admin"
	USERPASSWD = "123456"
	GINMODE = "release" // debug , release, test
	TCPSERVER = ":50001"
	BASEPATH = "./data/"
	WEEKFILE = BASEPATH + "Week/"
	TIMEFILE = BASEPATH + "time.json"
	SWITCHFILE = BASEPATH + "switch.json"
)