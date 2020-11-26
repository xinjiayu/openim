package web

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func Index(r *ghttp.Request) {

	communicateId := r.GetString("communicateId")
	userName := r.GetString("userName")

	apiUrl := g.Config().GetString("server.apiUrl")
	websocketServerAddr := g.Config().GetString("server.websocketServerAddr")

	data := g.Map{
		"communicateId":       communicateId,
		"userName":            userName,
		"apiUrl":              apiUrl,
		"websocketServerAddr": websocketServerAddr,
	}
	if err := r.Response.WriteTpl("index.html", data); err != nil {
		glog.Error(err.Error())
	}
}
