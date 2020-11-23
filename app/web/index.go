package web

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func Index(r *ghttp.Request) {

	communicateId := r.GetString("communicateId")
	userName := r.GetString("userName")

	data := g.Map{
		"communicateId": communicateId,
		"userName":      userName,
	}
	if err := r.Response.WriteTpl("index.html", data); err != nil {
		glog.Error(err.Error())
	}
}
