package models

type FeishuText struct {
	Webhook string `json:"webhook" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type FeishuSendText struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

type FeishuImage struct {
	Webhook string `json:"webhook" binding:"required"`
	Image   string `json:"image" binding:"required"`
}

type FeishuSendImage struct {
	MsgType string `json:"msg_type"`
	Content struct {
		ImageKey string `json:"image_key"`
	} `json:"content"`
}

type FeishuLink struct {
	Webhook  string `json:"webhook" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	LinkText string `json:"linkText" binding:"required"`
	LinkUrl  string `json:"linkUrl" binding:"required"`
}

type ActionsText struct {
	Tag  string `json:"tag"`
	Text struct {
		Content string `json:"content"`
		Tag     string `json:"tag"`
	} `json:"text"`
	Url   string   `json:"url"`
	Type  string   `json:"type"`
	Value struct{} `json:"value"`
}

type ElementsActions struct {
	Actions []*ActionsText `json:"actions"`
	Tag     string         `json:"tag"`
}

type ElementsText struct {
	Tag  string `json:"tag"`
	Text struct {
		Content string `json:"content"`
		Tag     string `json:"tag"`
	} `json:"text"`
}

type FeishuSendLink struct {
	MsgType string `json:"msg_type"`
	Card    struct {
		Config struct {
			WideScreenMode bool `json:"wide_screen_mode"`
			EnableForward  bool `json:"enable_forward"`
		} `json:"config"`
		Elements []interface{} `json:"elements"`
		Header   struct {
			Title struct {
				Content string `json:"content"`
				Tag     string `json:"tag"`
			} `json:"title"`
		} `json:"header"`
	} `json:"card"`
}
