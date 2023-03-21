package routers

import (
	"monitors-service-api-gitee/controllers/dingding"
	"monitors-service-api-gitee/middleware"

	"github.com/gin-gonic/gin"
)

func DingRouterInit(r *gin.Engine) {
	dingRouters := r.Group("/ding")
	{
		dingRouters.Use(middleware.Auth())
		dingRouters.POST("/sendLink", dingding.DingMessage{}.SendLink)
	}
}
