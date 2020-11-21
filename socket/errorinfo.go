package socket

import (
	"github.com/gogf/gf/errors/gerror"
)

var (
	// ErrorClose 连接关闭的错误信息
	ErrorClose = gerror.New("firetower is collapsed")
	// ErrorBlock block错误信息
	ErrorBlock = gerror.New("network congestion")
)
