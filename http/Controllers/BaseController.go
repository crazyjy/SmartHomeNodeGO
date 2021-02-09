package Controllers

import (
	"example.com/m/v2/base"
	"example.com/m/v2/config"
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