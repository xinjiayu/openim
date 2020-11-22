package main

import (
	"github.com/gogf/gf/frame/g"
	_ "github.com/mattn/go-sqlite3"
	"openim/app/service/manager"
	"openim/library/version"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	m := &manager.Manager{}
	go m.StartGrpcService(g.Config().GetString("topicServer.grpcPort"))
	m.StartSocketService(g.Config().GetString("topicServer.socketPort"))
}
