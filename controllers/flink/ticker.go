package flink

import (
	"fmt"
	"monitors-service-api-gitee/utils"
	"strconv"
	"time"

	"github.com/hucgqg/requests"
	"github.com/spf13/viper"
)

// 检查从/flink/addJobName传递过来的jobName是否再次传递，循环间隔时间由config.yaml/flink/timeTicker 值控制
// 检查到jobName未再次传递将进行电话告警，钉钉群告警
func TickerMonitor() {
	t, err := strconv.Atoi(viper.GetString("flink.timeTicker"))
	if err != nil {
		fmt.Println(err)
	}
	ticker := time.NewTicker(time.Second * time.Duration(t))
	defer ticker.Stop()
	for {
		<-ticker.C
		monitorJobList := CheckJobNameExist()
		for _, job := range monitorJobList {
			dingUrl := viper.GetString("localDingdingUrl")
			messageUrl := viper.GetString("flink.closeUrl") + *job
			title := fmt.Sprintf("Flink %v 服务告警", *job)
			msg := fmt.Sprintf("%v 服务出现错误，请及时处理，点击关闭本次报警", *job)
			// 电话报警
			utils.SendMessage(viper.GetString("flink.monitorUrl"), msg+messageUrl, title, msg)
			data := map[string]interface{}{
				"text":       msg,
				"title":      title,
				"messageUrl": messageUrl,
				"webhook":    viper.GetString("flink.monitorDingWebhook"),
				"picUrl":     "",
			}
			basicAuth := map[string]string{viper.GetString("username"): viper.GetString("password")}
			// 钉钉群报警
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
