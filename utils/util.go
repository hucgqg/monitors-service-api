package utils

import (
	"encoding/json"

	"github.com/hucgqg/requests"
)

// 将struct转成map[string]interface{}
func Struct2Map(b any) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	j, _ := json.Marshal(b)
	json.Unmarshal(j, &data)
	return data, nil
}

// 钉钉云监控电话报警
func SendMessage(url, ruleName, title, message string) {
	data := map[string]interface{}{
		"ruleName": ruleName,
		"title":    title,
		"message":  message,
	}

	r := requests.Request{
		Url:    url,
		Method: "POST",
		Data:   data,
	}
	if err := r.Body(); err != nil {
		panic(err)
	}
}
