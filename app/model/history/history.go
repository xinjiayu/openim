package history

import (
	"github.com/gogf/gf/errors/gerror"
)

//FindAll
func GetTopicAllData(topic string) ([]*Entity, error) {

	data, err := Model.FindAll("topic", topic)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, gerror.New("数据为空")

	}

	return data, nil

}
