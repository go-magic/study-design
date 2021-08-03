package factory

import (
	"github.com/go-magic/mid-server/task"
	"github.com/go-magic/study-design/server"
)

/*
IFactoryFunc 一般不单独使用,常常和其他设计模式结合使用,例如:观察者模式,抽象工厂等
*/
type IFactoryFunc func() task.Tasker

/*
NewHttpServer 工厂函数接口
*/
func NewHttpServer() task.Tasker {
	return &server.HttpServer{}
}
