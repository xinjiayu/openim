package gateway

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"time"

	"google.golang.org/grpc"
	pb "openim/grpc/manager"
	"openim/socket"
)

// buildManagerClient 实例化一个与topicManager连接的tcp链接
func buildManagerClient() {
	go func() {
	Retry:
		var err error
		grpcAddress := g.Config().GetString("websocketServer.grpcAddress")
		conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
		if err != nil {
			fmt.Println("[manager client] grpc connect error:", g.Config().GetString("websocketServer.grpcAddress"), err)
			time.Sleep(time.Duration(1) * time.Second)
			goto Retry
		}
		topicManageGrpc = pb.NewTopicServiceClient(conn)
		topicServiceAddr := g.Config().GetString("websocketServer.topicServiceAddr")
		topicManage = socket.NewClient(topicServiceAddr)

		topicManage.OnPush(func(sendMessage *socket.SendMessage) {
			TM.centralChan <- sendMessage
		})
		err = topicManage.Connect()
		if err != nil {
			glog.Info("[manager client] 等待主题管理器联机", g.Config().GetString("websocketServer.topicServiceAddr"))
			time.Sleep(time.Duration(1) * time.Second)
			goto Retry
		} else {
			glog.Info("[manager client] connected:", g.Config().GetString("websocketServer.topicServiceAddr"))
		}
	}()
}
