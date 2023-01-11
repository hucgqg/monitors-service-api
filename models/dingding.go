package models

type DingLink struct {
	MsgType string `json:"msgtype"`
	Link    struct {
		Webhook    string `json:"webhook" binding:"required"`
		Text       string `json:"text" binding:"required"`
		Title      string `json:"title" binding:"required"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl" binding:"required"`
	} `json:"link"`
}
