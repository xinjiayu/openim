package topicinfoService

import (
	"github.com/gogf/gf/os/glog"
	"openim/app/model/history"
	"openim/app/model/topicinfo"
)

func Add(topicInfoData *topicinfo.Entity) error {
	if _, err := topicinfo.Model.Replace(topicInfoData); err != nil {
		glog.Error(err.Error())
	}
	return nil
}

func GetTopicNewCount(topic, from string) int {
	ti, _ := topicinfo.Model.Where("topic", topic).Where("froma = ?", from).FindOne()
	glog.Info(ti)

	ht, _ := history.Model.Where("topic", topic).Where("froma != ?", from).Where("sendTime > ?", ti.Sendtime).Count()
	return ht
}
