package history

import (
	"github.com/gogf/gf/errors/gerror"
)

//GetTopicAllData
func GetTopicAllData(where ...interface{}) ([]*Entity, error) {

	data, err := Model.FindAll(where)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, gerror.New("数据为空")

	}

	return data, nil

}
