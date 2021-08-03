package server

import (
	"encoding/json"
	"github.com/go-magic/study-design/model"
	"net/http"
)

type HttpServer struct {
	subTask *model.HttpTask
}

/*
Check task.Tasker接口实例
*/
func (h *HttpServer) Check(subTask string) (string, error) {
	h.init()
	if err := h.parse(subTask); err != nil {
		return "", err
	}
	return h.check()
}

func (h *HttpServer) init() {
	h.subTask = model.NewHttpTask()
}

/*
反序列化实例
*/
func (h *HttpServer) parse(subTask string) error {
	return json.Unmarshal([]byte(subTask), h.subTask)
}

/*
检测实例
*/
func (h *HttpServer) check() (string, error) {
	res, err := h.request(h.subTask.Url)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	result := model.NewHttpResult(h.subTask)
	result.StatusCode = res.StatusCode
	b, mErr := json.Marshal(result)
	return string(b), mErr
}

/*
网页请求
*/
func (h *HttpServer) request(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	return client.Do(req)
}
