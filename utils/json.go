package utils

import (
	"encoding/json"
	"go-cli-starter/log"

	"go.uber.org/zap"
)

// JSONEncode 接受一个interface{}对象 返回 json字符串 []byte 格式
func JSONEncode(data interface{}) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Logger.Error(err.Error(), zap.String("service", "json"))
	}

	return jsonData
}

// JSONDecode json格式字符串 转换成 interface
func JSONDecode(jsonData []byte, data interface{}) interface{} {

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Logger.Error(err.Error(), zap.String("service", "json"))
	}
	return data
}
