package model

type HttpTask struct {
	Url string `json:"url"`
}

type HttpResult struct {
	StatusCode int    `json:"status_code"`
	CurrentUrl string `json:"current_url"`
}

func NewHttpTask() *HttpTask {
	return &HttpTask{}
}

func NewHttpResult(httpTask *HttpTask) *HttpResult {
	return &HttpResult{}
}
