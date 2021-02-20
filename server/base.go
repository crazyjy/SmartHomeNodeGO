package server

import (
	"smartHomeNode/v1/base"
	"smartHomeNode/v1/config"
)

type Base struct {
	base.Base
}

const (
	WEEKFILE = config.WEEKFILE
	TCPSERVER = config.TCPSERVER
)