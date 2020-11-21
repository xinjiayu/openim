package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	_ "openim/packed"
)

func init() {
	err := gtime.SetTimeZone("Asia/Shanghai") //设置系统时区
	if err != nil {
		glog.Error(err)
	}
	logPath := g.Config().GetString("logger.Path")
	err = glog.SetPath(logPath)
	if err != nil {
		glog.Error(err)
	}
}
