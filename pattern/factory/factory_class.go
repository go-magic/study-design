package factory

import (
	"github.com/go-magic/mid-server/task"
	"github.com/go-magic/study-design/server"
)

/*
IFactory 工厂接口
*/
type IFactory interface {
	CreateFactory() task.Tasker
}

/*
Factory 工厂类接口
*/
type Factory struct {
}

/*
CreateFactory 工厂类具体实现
*/
func (f Factory) CreateFactory() task.Tasker {
	return &server.HttpServer{}
}
