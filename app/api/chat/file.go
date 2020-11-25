package chat

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"openim/library/response"
)

// Upload uploads files to /tmp .
func Upload(r *ghttp.Request) {
	files := r.GetUploadFiles("upload-file")
	tmpdir := gconv.String(gtime.Now().Unix())
	newFileDir := "/upload/" + tmpdir + "/"
	tmpdir = g.Config().GetString("server.ServerRoot") + newFileDir
	names, err := files.Save(tmpdir)
	if err != nil {
		r.Response.WriteExit(err)
	}
	fileWebUrl := g.Config().GetString("server.WebUrl")
	fileWebUrl = fileWebUrl + newFileDir + names[0]
	response.JsonExit(r, 0, "上传成功", fileWebUrl)

}
