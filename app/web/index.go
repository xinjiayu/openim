package web

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {

	communicateId := r.GetString("communicateId")
	userName := r.GetString("userName")

	data := g.Map{
		"communicateId": communicateId,
		"userName":      userName,
	}
	r.Response.WriteTpl("index.html", data)
}
