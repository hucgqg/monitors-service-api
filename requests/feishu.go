package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"monitors-service-api/models"
	"net/http"
	"os"
	"path/filepath"
)

func GetTenantAccessToken() string {
	conf := models.Config{}
	url := conf.GetConfig().Service.Feishu.ApiUrl + "/open-apis/auth/v3/tenant_access_token/internal"
	data := map[string]string{
		"app_id":     conf.GetConfig().Service.Feishu.AppId,
		"app_secret": conf.GetConfig().Service.Feishu.AppSecret,
	}
	d, _ := json.Marshal(data)
	r := Request{Url: &url, Method: "POST", Data: &d, Headers: &map[string]string{}}
	msg, err := r.Body()
	if err != nil {
		fmt.Println(err)
	}
	var s = make(map[string]interface{})
	_ = json.Unmarshal(msg, &s)
	return s["tenant_access_token"].(string)
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
	client := &http.Client{}
	req, err := http.NewRequest("POST", uRL, payload)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Authorization", "Bearer "+GetTenantAccessToken())
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var uploadImage map[string]interface{}
	if err := json.Unmarshal(body, &uploadImage); err != nil {
		return "", err
	}
	if uploadImage["code"].(float64) != 0.0 {
		return uploadImage["msg"].(string), err
	}
	data := uploadImage["data"].(map[string]interface{})
	return data["image_key"].(string), nil
}
