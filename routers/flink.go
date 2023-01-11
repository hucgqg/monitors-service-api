package routers

import (
	"github.com/gin-gonic/gin"
	"monitors-service-api/controllers/flink"
	"monitors-service-api/middleware"
)

func FlinkRouterInit(r *gin.Engine) {
	flinkRouters := r.Group("/flink")
	{
		flinkRouters.GET("/delJobName/:jobName", flink.FlinkMontior{}.DeleteJobName)
		flinkRouters.Use(middleware.Auth())
		flinkRouters.GET("/checkJobName", flink.FlinkMontior{}.CheckJobNameExist)
		flinkRouters.POST("/addJobName", flink.FlinkMontior{}.AddJobName)
	}
}
