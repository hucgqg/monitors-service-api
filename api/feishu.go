package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"monitors-service-api-gitee/models"
	"os"
	"path/filepath"

	"github.com/hucgqg/requests"
)

func GetTenantAccessToken() string {
	conf := models.Config{}
	url := conf.GetConfig().Service.Feishu.ApiUrl + "/open-apis/auth/v3/tenant_access_token/internal"
	data := map[string]interface{}{
		"app_id":     conf.GetConfig().Service.Feishu.AppId,
		"app_secret": conf.GetConfig().Service.Feishu.AppSecret,
	}
	r := requests.Request{
		Url:    url,
		Method: "POST",
		Data:   data,
	}
	r.Body()
	return r.RepInfo["tenant_access_token"].(string)
}

func GetImageKey(imagePath string) (string, error) {
	conf := models.Config{}
	uRL := conf.GetConfig().Service.Feishu.ApiUrl + "/open-apis/im/v1/images"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("image_type", "message")
	file, errFile4 := os.Open(imagePath)
	defer file.Close()
	part4, errFile4 := writer.CreateFormFile("image", filepath.Base(imagePath))
	_, errFile4 = io.Copy(part4, file)
	if errFile4 != nil {
		return "", errFile4
	}
	if err := writer.Close(); err != nil {
		return "", err
	}
	d, _ := json.Marshal(payload)
	data := make(map[string]interface{})
	if err := json.Unmarshal(d, data); err != nil {
		fmt.Println(err)
	}
	headersAdd, headersSet := make(map[string]string), make(map[string]string)
	headersAdd["Content-Type"] = "multipart/form-data"
	headersAdd["Authorization"] = "Bearer " + GetTenantAccessToken()
	headersSet["Content-Type"] = writer.FormDataContentType()
	r := requests.Request{
		Data:       data,
		Method:     "POST",
		Url:        uRL,
		HeadersAdd: headersAdd,
		HeadersSet: headersSet,
	}
	r.Body()
	if r.RepInfo["code"].(float64) != 0.0 {
		return r.RepInfo["msg"].(string), fmt.Errorf("get image key faild")
	}
	respData := r.RepInfo["data"].(map[string]interface{})
	return respData["image_key"].(string), nil
}
