package routers

import (
	"monitors-service-api-gitee/controllers/feishu"
	"monitors-service-api-gitee/middleware"

	"github.com/gin-gonic/gin"
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
