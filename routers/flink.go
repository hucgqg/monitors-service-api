package routers

import (
	"monitors-service-api-gitee/controllers/flink"
	"monitors-service-api-gitee/middleware"

	"github.com/gin-gonic/gin"
)

func FlinkRouterInit(r *gin.Engine) {
	flinkRouters := r.Group("/flink")
	{
		flinkRouters.GET("/delJobName/:jobName", flink.FlinkMontior{}.DeleteJobName)
		flinkRouters.Use(middleware.Auth())
		flinkRouters.POST("/addJobName", flink.FlinkMontior{}.AddJobName)
	}
}
