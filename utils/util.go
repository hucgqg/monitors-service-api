package utils

import (
	"fmt"
	"reflect"

	"github.com/hucgqg/requests"
)

// 将struct转成map[string]interface{}
func Struct2Map(b any) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	objT := reflect.TypeOf(b)

	fmt.Println("------------ objT", objT)
	fmt.Println("------------ objT.NumField()", objT.NumField())
	objV := reflect.ValueOf(b)
	fmt.Println("------------ objV", objV)

	for i := 0; i < objT.NumField(); i++ {
		fileName, ok := objT.Field(i).Tag.Lookup("json")
		if ok {

			fmt.Println("------------ objV.Field(i).Interface()", objV.Field(i).Interface())

			data[fileName] = objV.Field(i).Interface()
		} else {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
	}
	return data, nil
}

// 钉钉云监控电话报警
func SendMessage(url, ruleName, title, message string) {
	data := map[string]interface{}{
		"ruleName": ruleName,
		"title":    title,
		"message":  message,
	}
	request := requests.Request{
		Url:       &url,
		Method:    "POST",
		Data:      &data,
		Headers:   &map[string]string{},
		BasicAuth: &map[string]string{},
	}
	if err := request.Body(); err != nil {
		panic(err)
	}
}
