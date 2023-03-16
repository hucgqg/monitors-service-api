package models

type RespData struct {
	Status bool      `json:"status"`
	Data   []*string `json:"data"`
	Msg    string    `json:"msg"`
	T      string    `json:"t"`
}
