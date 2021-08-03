package decorator

import (
	"errors"
)

type IDecoratorFunc func(interface{}) error

func sendGateway(data interface{}) error {
	return errors.New("发送失败")
}

func addData(data interface{}) error {
	return errors.New("发送失败")
}

/*
TryFunc 装饰器函数(当有多个函数都需要进行多次尝试,可以使用装饰器函数进行多次尝试)
*/
func TryFunc(decorator IDecoratorFunc, data interface{}, times int) error {
	var err = errors.New("times=0")
	for i := 0; i < times; i++ {
		if err = decorator(data); err == nil {
			return nil
		}
	}
	return err
}

/*
TryDecorator 装饰器函数(当有多个函数都需要进行多次尝试,可以使用装饰器函数进行多次尝试)
*/
func TryDecorator(decorator IDecorator, data interface{}, times int) error {
	var err = errors.New("times=0")
	for i := 0; i < times; i++ {
		if err = decorator.Execute(data); err == nil {
			return nil
		}
	}
	return err
}
