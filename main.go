package main

import (
	"github.com/gin-gonic/gin"
	"monitors-service-api/routers"
)

func main() {
	// 生产环境模式
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	routers.FeishuRouterInit(r)
	routers.DingRouterInit(r)
	routers.FlinkRouterInit(r)

	err := r.Run()
	if err != nil {
		return
	}
}
