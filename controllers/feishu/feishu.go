package feishu

import (
	"monitors-service-api-gitee/api"
	"monitors-service-api-gitee/models"
	"monitors-service-api-gitee/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hucgqg/requests"
)

type FeishuMessage struct{}

func (f FeishuMessage) SendText(c *gin.Context) {
	text := models.FeishuText{}
	sendText := models.FeishuSendText{}
	resp := models.RespData{T: time.Now().Format("2006-01-02 15:04:05:000")}
	if err := c.BindJSON(&text); err != nil {
		resp.Msg = "请求失败:" + err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	sendText.MsgType = "text"
	sendText.Content.Text = text.Content
	data, err := utils.Struct2Map(sendText)
	if err != nil {
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	r := requests.Request{
		Url:    text.Webhook,
		Method: "POST",
		Data:   data,
	}
	r.Body()
	if r.RepInfo["StatusCode"] == 0.0 {
		resp.Msg = "请求成功"
		resp.Status = true
		c.JSON(http.StatusOK, &resp)
	} else {
		resp.Msg = r.RepInfo["msg"].(string)
		c.JSON(http.StatusOK, &resp)
	}
}

func (f FeishuMessage) SendImage(c *gin.Context) {
	sendImage := models.FeishuSendImage{}
	resp := models.RespData{T: time.Now().Format("2006-01-02 15:04:05:000")}
	webhook := c.PostForm("webhook")
	file, _ := c.FormFile("image")
	if webhook == "" || file == nil {
		resp.Msg = "请求失败:webhook和需要上传文件不能为空"
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	dst := "./upload/" + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		resp.Msg = "请求失败:" + err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	imageKey, err := api.GetImageKey(dst)
	if err != nil {
		resp.Msg = "请求失败:" + err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	sendImage.Content.ImageKey = imageKey
	sendImage.MsgType = "image"
	data, _ := utils.Struct2Map(sendImage)
	r := requests.Request{
		Url:    webhook,
		Method: "POST",
		Data:   data,
	}
	r.Body()
	if r.RepInfo["StatusCode"] != 0.0 {
		resp.Msg = "请求失败:" + r.RepInfo["msg"].(string)
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	_ = os.Remove(dst)
	resp.Msg = "请求成功"
	resp.Status = true
	c.JSON(http.StatusOK, &resp)
}

func (f FeishuMessage) SendLink(c *gin.Context) {
	resp := models.RespData{T: time.Now().Format("2006-01-02 15:04:05:000")}
	linkData := models.FeishuLink{}
	if err := c.BindJSON(&linkData); err != nil {
		resp.Msg = "请求失败: " + err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	actionsText := models.ActionsText{}
	actionsText.Tag = "button"
	actionsText.Text.Tag = "lark_md"
	actionsText.Text.Content = linkData.LinkText
	actionsText.Url = linkData.LinkUrl

	elementsText := models.ElementsText{}
	elementsText.Tag = "div"
	elementsText.Text.Tag = "lark_md"
	elementsText.Text.Content = linkData.Text

	elementsActions := models.ElementsActions{}
	elementsActions.Tag = "action"
	elementsActions.Actions = append(elementsActions.Actions, &actionsText)

	sendLink := models.FeishuSendLink{}
	sendLink.MsgType = "interactive"
	sendLink.Card.Config.WideScreenMode = true
	sendLink.Card.Config.EnableForward = true
	sendLink.Card.Header.Title.Tag = "plain_text"
	sendLink.Card.Header.Title.Content = linkData.Title
	sendLink.Card.Elements = append(sendLink.Card.Elements, elementsText)
	sendLink.Card.Elements = append(sendLink.Card.Elements, elementsActions)

	data, _ := utils.Struct2Map(sendLink)
	r := requests.Request{
		Url:    linkData.Webhook,
		Method: "POST",
		Data:   data,
	}
	r.Body()
	if r.RepInfo["StatusCode"] != 0.0 {
		resp.Msg = "请求失败: " + r.RepInfo["msg"].(string)
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	resp.Msg = "请求成功"
	resp.Status = true
	c.JSON(http.StatusOK, &resp)
}
