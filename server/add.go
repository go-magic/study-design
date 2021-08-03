package server

import (
	"encoding/json"
	"github.com/go-magic/study-design/model"
)

type Add struct {
	subTask *model.AddTask
}

/*
Check task.Tasker接口实例
*/
func (a *Add) Check(subTask string) (string, error) {
	if err := a.parse(subTask); err != nil {
		return "", err
	}
	return a.check()
}

/*
反序列化实例
*/
func (a *Add) parse(subTask string) error {
	return json.Unmarshal([]byte(subTask), a.subTask)
}

/*
检测实例
*/
func (a *Add) check() (string, error) {
	var sum int
	for _, arg := range a.subTask.Args {
		sum += arg
	}
	b, err := json.Marshal(model.AddResult{Sum: sum})
	if err != nil {
		return "", err
	}
	return string(b), err
}
