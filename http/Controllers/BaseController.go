package Controllers

import (
	"smartHomeNode/v1/base"
	"smartHomeNode/v1/config"
)

const (
	USERMAME = config.USERMAME
	USERPASSWD = config.USERPASSWD
	TIMEFILE = config.TIMEFILE
	WEEKFILE = config.WEEKFILE
	SWITCHFILE = config.SWITCHFILE
)

type BaseController struct {
	base.Base
}