package routers

import (
	"github.com/gin-gonic/gin"
	"monitors-service-api/controllers/feishu"
	"monitors-service-api/middleware"
)

func FeishuRouterInit(r *gin.Engine) {
	feishuRouters := r.Group("/feishu")
	{
		feishuRouters.Use(middleware.Auth())
		feishuRouters.POST("/sendText", feishu.FeishuMessage{}.SendText)
		feishuRouters.POST("/sendImage", feishu.FeishuMessage{}.SendImage)
		feishuRouters.POST("/sendLink", feishu.FeishuMessage{}.SendLink)
	}
}
