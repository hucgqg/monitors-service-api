package models

type FlinkData struct {
	JobName string `json:"jobName" form:"jobName" binding:"required"`
}
