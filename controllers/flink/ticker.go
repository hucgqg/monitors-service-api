package flink

import (
	"fmt"
	"time"

	"gitee.com/hcqcode/requests"
)

var (
	flinkMonitorUrl string = ""
	flinkWebhook    string = ""
)

func SendMessage(url, ruleName, title, message string) {
	data := map[string]interface{}{
		"ruleName": ruleName,
		"title":    title,
		"message":  message,
	}
	request := requests.Request{Url: &url, Method: "POST", Data: &data, Headers: &map[string]string{}, BasicAuth: &map[string]string{}}
	request.Body()
}

func TickerMonitor() {
	ticker := time.NewTicker(time.Minute * 3)
	defer ticker.Stop()
	for {
		<-ticker.C
		monitorJobList := CheckJobNameExist()
		for _, job := range monitorJobList {
			dingUrl := ""
			messageUrl := fmt.Sprintf("https://test.com/%v", *job)
			title := fmt.Sprintf("Flink %v 服务告警", *job)
			msg := fmt.Sprintf("%v 服务出现错误，请及时处理，点击关闭本次报警", *job)
			SendMessage(flinkMonitorUrl, msg+messageUrl, title, msg)
			data := map[string]interface{}{
				"text":       msg,
				"title":      title,
				"messageUrl": messageUrl,
				"webhook":    flinkWebhook,
				"picUrl":     "",
			}
			basicAuth := map[string]string{"admin": "M473f5eef2d5ced_"}
			request := requests.Request{Url: &dingUrl, Method: "POST", Headers: &map[string]string{}, Data: &data, BasicAuth: &basicAuth}
			request.Body()
		}
	}
}
