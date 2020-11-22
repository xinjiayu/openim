package historyService

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"openim/app/model/history"
)

func Add(historyData *history.Entity) error {

	historyData.Sendtime = gconv.Int(gtime.Now().Unix())
	history.Model.Insert(historyData)
	return nil
}

func GetDataBeyTopic(topic string) []*history.Entity {
	data, _ := history.GetTopicAllData(topic)
	return data
}
