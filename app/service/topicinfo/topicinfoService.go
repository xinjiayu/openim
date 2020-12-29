package topicinfoService

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"openim/app/model/topicinfo"
)

func Add(topicInfoData *topicinfo.Entity) error {

	if _, err := topicinfo.Model.Replace(topicInfoData); err != nil {
		glog.Error(err.Error())
	}
	return nil
}

//SetReadState
func SetReadState(topic, from string) error {
	tlist, _ := topicinfo.Model.Where("topic", topic).Where("froma != ?", from).FindAll()
	for _, v := range tlist {
		//更新topic信息状态
		topicInfo := new(topicinfo.Entity)
		topicInfo.Topic = v.Topic
		topicInfo.Froma = v.Froma
		topicInfo.Num = 0
		topicInfo.Sendtime = gconv.Int(gtime.Now().Unix())
		if topicInfo.Froma != "" {
			if err := Add(topicInfo); err != nil {
				glog.Error(err.Error())
			}
		}
	}

	return nil

}

//func GetTopicNewCount(topic, from string) int {
//	ti, _ := topicinfo.Model.Where("topic", topic).Where("froma = ?", from).FindOne()
//	return ti.Num
//}

//GetTopicByFrom 获取指定用户是否有最新未读记录
func GetTopicByFrom(from string) []*topicinfo.Entity {

	ti, _ := topicinfo.Model.Where("froma = ?", from).FindAll()

	var res []*topicinfo.Entity
	for _, t := range ti {
		topiclist, _ := topicinfo.Model.Where("topic = ?", t.Topic).FindAll()
		for _, topic := range topiclist {
			if topic.Froma != from {
				res = append(res, topic)
			}
		}

	}

	return res
}
