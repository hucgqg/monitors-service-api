package flink

import (
	"fmt"
	"monitors-service-api-gitee/utils"
	"time"

	"github.com/hucgqg/requests"
)

var (
	flinkMonitorUrl string = ""
	flinkWebhook    string = ""
)

// 每三分钟检查从/flink/addJobName传递过来的jobName是否再次传递
// 检查到jobName未再次传递将进行电话告警，钉钉群告警
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
			utils.SendMessage(flinkMonitorUrl, msg+messageUrl, title, msg)
			data := map[string]interface{}{
				"text":       msg,
				"title":      title,
				"messageUrl": messageUrl,
				"webhook":    flinkWebhook,
				"picUrl":     "",
			}
			basicAuth := map[string]string{"admin": "M473f5eef2d5ced_"}
			r := requests.Request{
				Url:       dingUrl,
				Method:    "POST",
				Data:      data,
				BasicAuth: basicAuth,
			}
			r.Body()
		}
	}
}
