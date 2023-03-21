package flink

import (
	"fmt"
	"monitors-service-api-gitee/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	jobList        []*string
	checkJobList   []*string
	monitorJobList []*string
)

func checkJobName(jobName string, jobNameList []*string) bool {
	for _, v := range jobNameList {
		if *v == jobName {
			return true
		}
	}
	return false
}

func CheckJobNameExist() []*string {
	monitorJobList = monitorJobList[:0]
	for _, v := range checkJobList {
		if !checkJobName(*v, jobList) {
			monitorJobList = append(monitorJobList, v)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "	需要报警，job_name:	", *v)
		}
	}
	jobList = jobList[:0]
	return monitorJobList
}

type FlinkMontior struct{}

func (f FlinkMontior) AddJobName(c *gin.Context) {
	resp := models.RespData{T: time.Now().Format("2006-01-02 15:04:05:000")}
	flink := models.FlinkData{}
	if err := c.BindJSON(&flink); err != nil {
		resp.Msg = "请求失败: " + err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}
	if !checkJobName(flink.JobName, jobList) {
		jobList = append(jobList, &flink.JobName)
	}
	if !checkJobName(flink.JobName, checkJobList) {
		fmt.Println(resp.T, "	添加job_name:	", flink.JobName)
		checkJobList = append(checkJobList, &flink.JobName)
	}
	for i, v := range checkJobList {
		fmt.Printf("%v		第%d个job_name:	%v\n", resp.T, i+1, *v)
	}
	resp.Msg = "请求成功"
	resp.Status = true
	resp.Data = checkJobList
	c.JSON(http.StatusOK, &resp)

}

func (f FlinkMontior) DeleteJobName(c *gin.Context) {
	jobName := c.Param("jobName")
	if checkJobName(jobName, checkJobList) {
		for i, v := range checkJobList {
			if *v == jobName {
				checkJobList = append(checkJobList[:i], checkJobList[i+1:]...)
			}
		}
		c.String(http.StatusOK, "关闭 %s 告警成功", jobName)
		return
	}
	c.String(http.StatusOK, "报警已关闭")
}
