package decorator

import "errors"

type IDecorator interface {
	Execute(data interface{}) error
}

type Gateway struct {
}

func (g Gateway) Execute(data interface{}) error {
	return errors.New("发送失败")
}
