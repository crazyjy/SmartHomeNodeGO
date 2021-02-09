package server

import (
	"example.com/m/v2/base"
	"example.com/m/v2/config"
)

type Base struct {
	base.Base
}

const (
	WEEKFILE = config.WEEKFILE
	TCPSERVER = config.TCPSERVER
)