package historyService

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"openim/app/model/history"
	"openim/app/model/topicinfo"
	topicinfoService "openim/app/service/topicinfo"
)

func Add(historyData *history.Entity) error {

	historyData.Sendtime = gconv.Int(gtime.Now().Unix())
	history.Model.Insert(historyData)

	//更新topic信息状态
	topicInfo := new(topicinfo.Entity)
	topicInfo.Topic = historyData.Topic
	topicInfo.Froma = historyData.Froma
	topicInfo.Num = 0
	topicInfo.Sendtime = historyData.Sendtime
	if topicInfo.Froma != "" {
		topicinfoService.Add(topicInfo)
	}

	return nil
}

func GetDataBeyTopic(topic, from string) []*history.Entity {
	data, _ := history.GetTopicAllData("topic", topic)

	//更新topic信息状态
	topicInfo := new(topicinfo.Entity)
	topicInfo.Topic = topic
	topicInfo.Froma = from
	topicInfo.Num = 0
	topicInfo.Sendtime = gconv.Int(gtime.Now().Unix())
	if topicInfo.Froma != "" {
		topicinfoService.Add(topicInfo)
	}

	return data
}
