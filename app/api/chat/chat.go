package chat

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gorilla/websocket"
	"net/http"
	"openim/app/service/gateway"
	historyService "openim/app/service/history"
	topicinfoService "openim/app/service/topicinfo"
	"openim/library/algorithm/snowFlake"
	"openim/library/response"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type messageInfo struct {
	From string `json:"from"`
	Data string `json:"data"`
	Type string `json:"type"`
}

// GlobalIdWorker 全局唯一id生成器
var GlobalIdWorker *snowFlake.Worker

func init() {
	// 全局唯一id生成器
	GlobalIdWorker, _ = snowFlake.NewWorker(1)
	// 如果是集群环境  一定一定要给每个服务设置唯一的id
	// 取值范围 1-1024
	gateway.ClusterId = 1
	gateway.Init()
}

// Websocket http转websocket连接 并实例化firetower
//func Websocket(w http.ResponseWriter, r *http.Request) {
func Websocket(r *ghttp.Request) {

	// 做用户身份验证

	// 验证成功才升级连接
	ws, _ := upgrader.Upgrade(r.Response.Writer, r.Request, nil)

	id := GlobalIdWorker.GetId()
	tower := gateway.BuildTower(ws, strconv.FormatInt(id, 10))

	tower.SetReadHandler(func(fire *gateway.FireInfo) bool {
		// 做发送验证
		// 判断发送方是否有权限向到达方发送内容
		if err := tower.Publish(fire); err != nil {
			glog.Error(err.Error())
		}
		return true
	})

	tower.SetReadTimeoutHandler(func(fire *gateway.FireInfo) {
		messageInfo := new(messageInfo)
		err := json.Unmarshal(fire.Message.Data, messageInfo)
		if err != nil {
			return
		}
		messageInfo.Type = "timeout"
		b, _ := json.Marshal(messageInfo)
		err = tower.ToSelf(b)
		if err != gateway.ErrorClose {
			fmt.Println("err:", err)
		}
	})

	tower.SetBeforeSubscribeHandler(func(context *gateway.FireLife, topic []string) bool {
		// 这里用来判断当前用户是否允许订阅该topic

		return true
	})

	tower.SetSubscribeHandler(func(context *gateway.FireLife, topic []string) bool {
		for _, v := range topic {
			num := tower.GetConnectNum(v)
			// 继承订阅消息的context
			var pushmsg = gateway.NewFireInfo(tower, context)
			glog.Info("============", pushmsg)
			pushmsg.Message.Topic = v
			pushmsg.Message.Data = []byte(fmt.Sprintf("{\"type\":\"onSubscribe\",\"data\":%d}", num))
			if err := tower.Publish(pushmsg); err != nil {
				glog.Error(err.Error())
			}
		}
		return true
	})

	tower.SetUnSubscribeHandler(func(context *gateway.FireLife, topic []string) bool {
		for _, v := range topic {
			num := tower.GetConnectNum(v)
			var pushmsg = gateway.NewFireInfo(tower, context)
			pushmsg.Message.Topic = v
			pushmsg.Message.Data = []byte(fmt.Sprintf("{\"type\":\"onUnsubscribe\",\"data\":%d}", num))
			if err := tower.Publish(pushmsg); err != nil {
				glog.Error(err.Error())
			}
		}
		return true
	})

	tower.Run()
}

func GetHistory(r *ghttp.Request) {
	topic := r.GetString("topic")
	from := r.GetString("from")

	data := historyService.GetDataBeyTopic(topic, from)
	response.JsonExit(r, 0, "历史数据", data)

}

//获取指定from的会话ID列表
func GetTopicByFrom(r *ghttp.Request) {
	from := r.GetString("from")
	glog.Info(from)

	data := topicinfoService.GetTopicByFrom(from)
	response.JsonExit(r, 0, "新信息数", data)
}
