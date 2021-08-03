package strategy

import "github.com/go-magic/mid-server/task"

type IStrategy interface {
	Check(sunTask string) (string, error)
}

type Strategy struct {
	tasker task.Tasker
}

func (s Strategy) Check(sunTask string) (string, error) {
	return sunTask, nil
}

/*
结合观察者、工厂模式使用,参考server、observe、factory模块
*/

func NewStrategy(tasker task.Tasker) IStrategy {
	return Strategy{tasker: tasker}
}
