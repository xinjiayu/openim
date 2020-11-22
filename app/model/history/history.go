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

//GetTopicCount
func GetTopicCount(where ...interface{}) (int, error) {

	data, err := Model.Count(where)
	if err != nil {
		return 0, err
	}
	return data, nil
}
