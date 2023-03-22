package models

type RespData struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"message"`
	T      string      `json:"t"`
}
