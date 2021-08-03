package observe

import (
	"github.com/go-magic/mid-server/task"
	"github.com/go-magic/study-design/pattern/factory"
)

type IObserve interface {
	Register(int, factory.IFactory)
	GetTasker(int) task.Tasker
	Exit()
}

type IObserveFunc interface {
	Register(int, factory.IFactoryFunc)
	GetTasker(int) task.Tasker
	Exit()
}
