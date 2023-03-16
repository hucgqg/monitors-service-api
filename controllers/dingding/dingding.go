package dingding

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"monitors-service-api/models"
	"monitors-service-api/requests"
	"net/http"
	"time"
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
	data, _ := json.Marshal(link)
	r := requests.Request{Data: &data, Method: "POST", Url: &link.Link.Webhook, Headers: &map[string]string{}}
	msg, err := r.Body()
	if err != nil {
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	var respData map[string]interface{}
	_ = json.Unmarshal(msg, &respData)
	if respData["errcode"] != 0.0 {
		resp.Msg = respData["errmsg"].(string)
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	resp.Status = true
	resp.Msg = "请求成功"
	c.JSON(http.StatusOK, &resp)
}
