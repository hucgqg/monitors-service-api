package routers

import (
	"github.com/gin-gonic/gin"
	"monitors-service-api/controllers/dingding"
	"monitors-service-api/middleware"
)

func DingRouterInit(r *gin.Engine) {
	dingRouters := r.Group("/ding")
	{
		dingRouters.Use(middleware.Auth())
		dingRouters.POST("/sendLink", dingding.DingMessage{}.SendLink)
	}
}
