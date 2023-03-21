package dingding

import (
	"monitors-service-api-gitee/models"
	"monitors-service-api-gitee/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hucgqg/requests"
)

type DingMessage struct{}

func (d DingMessage) SendLink(c *gin.Context) {
	resp := models.RespData{T: time.Now().Format("2006-01-02 15:04:05:000")}
	link := models.DingLink{}
	link.MsgType = "link"
	if err := c.BindJSON(&link.Link); err != nil {
		resp.Msg = "请求失败: " + err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	data, err := utils.Struct2Map(link)
	if err != nil {
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	r := requests.Request{
		Data:   data,
		Method: "POST",
		Url:    link.Link.Webhook,
	}
	r.Body()
	if r.RepInfo["errcode"] != 0.0 {
		resp.Msg = r.RepInfo["errmsg"].(string)
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	resp.Status = true
	resp.Msg = "请求成功"
	c.JSON(http.StatusOK, &resp)
}
