package dingding

import (
	"encoding/json"
	"fmt"
	"monitors-service-api/models"
	"monitors-service-api/utils"
	"net/http"
	"time"

	"github.com/hucgqg/requests"

	"github.com/gin-gonic/gin"
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
	l, _ := json.Marshal(&link)
	data, err := utils.Struct2Map(l)
	fmt.Println("---- dingding  data  Struct2Map", data)
	if err != nil {
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	json.Unmarshal(l, &d)
	r := requests.Request{Data: &data, Method: "POST", Url: &link.Link.Webhook, Headers: &map[string]string{}}
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
