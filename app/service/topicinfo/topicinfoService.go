package topicinfoService

import (
	"github.com/gogf/gf/os/glog"
	"openim/app/model/history"
	"openim/app/model/topicinfo"
)

func Add(topicInfoData *topicinfo.Entity) error {
	topicinfo.Model.Replace(topicInfoData)
	return nil
}

func GetTopicNewCount(topic, from string) int {
	ti, _ := topicinfo.Model.Where("topic", topic).Where("froma = ?", from).FindOne()
	glog.Info(ti)

	ht, _ := history.Model.Where("topic", topic).Where("froma != ?", from).Where("sendTime > ?", ti.Sendtime).Count()
	return ht
}
