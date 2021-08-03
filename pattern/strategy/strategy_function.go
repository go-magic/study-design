package strategy

import "github.com/go-magic/mid-server/task"

func IStrategyFunc(tasker task.Tasker, subTask string) (string, error) {
	return tasker.Check(subTask)
}
