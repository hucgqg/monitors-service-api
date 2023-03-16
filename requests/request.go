package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Url     *string
	Method  string
	Headers *map[string]string
	Data    *[]byte
}

func (r *Request) Body() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, *r.Url, bytes.NewBuffer(*r.Data))
	if err != nil {
		return *r.Data, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	for k, v := range *r.Headers {
		req.Header.Add(k, v)
	}
	rep, err := client.Do(req)
	if err != nil {
		return *r.Data, err
	}
	defer rep.Body.Close()
	body, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return *r.Data, err
	}
	return body, nil
}
