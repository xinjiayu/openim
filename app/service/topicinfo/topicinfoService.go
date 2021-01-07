package topicinfoService

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"openim/app/model/topicinfo"
	"openim/library/tools"
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

//GetTopicByFrom 获取指定用户是否有最新未读记录
func GetTopicByFrom(from string) []*topicinfo.Entity {

	topicInfoList := make(map[string]*topicinfo.Entity)
	checkNewUserTopic(topicInfoList, from)

	var res []*topicinfo.Entity
	for _, tl := range topicInfoList {
		res = append(res, tl)
	}
	return res
}

//检查
func checkNewUserTopic(res map[string]*topicinfo.Entity, from string) {

	ti, _ := topicinfo.Model.FindAll()
	for _, t := range ti {
		if t.Froma != from {
			commId := tools.CreateCommunicateId(gconv.Int(from), gconv.Int(t.Froma))
			getTopicInfoList(res, from, commId)

		}

	}
}

//getTopicInfoList 获取会话列表
func getTopicInfoList(res map[string]*topicinfo.Entity, from, topic string) {
	topiclist, _ := topicinfo.Model.Where("topic = ?", topic).FindAll()
	for _, t := range topiclist {
		if t.Froma != from {
			res[t.Topic] = t
		}
	}
}
