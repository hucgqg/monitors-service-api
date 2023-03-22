package main

import (
	"monitors-service-api-gitee/controllers/flink"
	"monitors-service-api-gitee/models"
	"monitors-service-api-gitee/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 生产环境模式
	gin.SetMode(gin.ReleaseMode)
	// 初始化配置文件
	models.InitConfig()

	r := gin.Default()
	routers.FeishuRouterInit(r)
	routers.DingRouterInit(r)
	routers.FlinkRouterInit(r)

	go func() {
		flink.TickerMonitor()
	}()

	err := r.Run()
	if err != nil {
		return
	}
}
