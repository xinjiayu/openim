package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"net/http"
	"openim/library/tools"
)

func White(r *ghttp.Request) {
	accessOk := true
	//获取客户端IP
	cip := r.GetRemoteIp()
	ipArr := g.Config().GetArray("system.whitelist")
	if len(ipArr) > 0 {
		accessOk = tools.IsContains(cip, ipArr)
	}

	if accessOk {
		r.Middleware.Next()
	} else {
		glog.Info(cip, "此IP无权访问服务")
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
