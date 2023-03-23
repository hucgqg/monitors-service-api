package main

import (
	"monitors-service-api-gitee/controllers/flink"
	"monitors-service-api-gitee/models"
	"monitors-service-api-gitee/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 生产环境模式
	gin.SetMode(gin.ReleaseMode)
	// 初始化配置文件
	models.InitConfig()

	// 路由注册
	r := gin.Default()
	routers.FeishuRouterInit(r)
	routers.DingRouterInit(r)
	routers.FlinkRouterInit(r)

	// 启动Flink定时监控
	go func() {
		flink.TickerMonitor()
	}()

	// 启动时端口设置
	if viper.GetString("port") == "" {
		r.Run()
	} else {
		r.Run(":" + viper.GetString("port"))
	}
}
