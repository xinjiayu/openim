package main

import (
	"github.com/gogf/gf/frame/g"
	"openim/app/service/manager"
)

func main() {
	m := &manager.Manager{}
	go m.StartGrpcService(g.Config().GetString("topicServer.grpcPort"))
	m.StartSocketService(g.Config().GetString("topicServer.socketPort"))
}
