package observe

import (
	"github.com/go-magic/mid-server/task"
	"github.com/go-magic/study-design/pattern/factory"
	"sync"
)

/*
FactoryFuncRegister 注册结构
*/
type FactoryFuncRegister struct {
	Factory factory.IFactoryFunc
	MsgId   int
}

type factoryFuncObserve struct {
	exit           chan struct{}
	registerChan   chan FactoryFuncRegister
	registerCenter map[int]factory.IFactoryFunc
	getTaskerChan  chan int
	sendTaskerChan chan task.Tasker
}

var (
	factoryFuncOb *factoryFuncObserve
	obOnce        sync.Once
)

func newFactoryFuncObserve() *factoryFuncObserve {
	return &factoryFuncObserve{
		exit:           make(chan struct{}),
		registerChan:   make(chan FactoryFuncRegister),
		registerCenter: make(map[int]factory.IFactoryFunc),
		getTaskerChan:  make(chan int),
	}
}

/*
GetFactoryFuncObserveInstance observe实例
*/
func GetFactoryFuncObserveInstance() *factoryFuncObserve {
	obOnce.Do(func() {
		factoryFuncOb = newFactoryFuncObserve()
		go factoryFuncOb.start()
	})
	return factoryFuncOb
}

/*
事件转发
*/
func (o factoryFuncObserve) start() {
	for {
		select {
		case <-o.exit:
			return
		case r := <-o.registerChan:
			o.register(r)
		case msgId := <-o.getTaskerChan:
			o.sendTasker(msgId)
		}
	}
}

/*
Register 消息注册外部接口
*/
func (o factoryFuncObserve) Register(msgId int, f factory.IFactoryFunc) {
	o.registerChan <- FactoryFuncRegister{MsgId: msgId, Factory: f}
}

/*
消息注册内部接口,去除锁
*/
func (o factoryFuncObserve) register(r FactoryFuncRegister) {
	o.registerCenter[r.MsgId] = r.Factory
}

/*
GetTasker 获取消息外部接口
*/
func (o factoryFuncObserve) GetTasker(msgId int) task.Tasker {
	o.getTaskerChan <- msgId
	return <-o.sendTaskerChan
}

/*
获取tasker,通过factory生产tasker
*/
func (o factoryFuncObserve) sendTasker(msgId int) {
	f, exist := o.registerCenter[msgId]
	if !exist {
		o.sendTaskerChan <- nil
		return
	}
	o.sendTaskerChan <- f()
}

/*
Exit 退出
*/
func (o factoryFuncObserve) Exit() {
	go func() {
		o.exit <- struct{}{}
	}()
}
