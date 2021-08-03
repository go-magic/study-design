package observe

import (
	"github.com/go-magic/mid-server/task"
	"github.com/go-magic/study-design/pattern/factory"
	"sync"
)

/*
观察者、工厂、策略模式
*/

/*
FactoryRegister 注册结构
*/
type FactoryRegister struct {
	Factory factory.IFactory
	MsgId   int
}

type factoryObserve struct {
	exit           chan struct{}            /*退出*/
	registerChan   chan FactoryRegister     /*注册*/
	registerCenter map[int]factory.IFactory /*注册中心*/
	getTaskerChan  chan int                 /*根据msgId策略选择执行任务实体*/
	sendTaskerChan chan task.Tasker         /*发送tasker*/
}

var (
	factoryOb *factoryObserve
	once      sync.Once
)

func newFactoryObserve() *factoryObserve {
	return &factoryObserve{
		exit:           make(chan struct{}),
		registerChan:   make(chan FactoryRegister),
		registerCenter: make(map[int]factory.IFactory),
		getTaskerChan:  make(chan int),
	}
}

/*
GetFactoryObserveInstance observe实例
*/
func GetFactoryObserveInstance() *factoryObserve {
	once.Do(func() {
		factoryOb = newFactoryObserve()
		go factoryOb.start()
	})
	return factoryOb
}

/*
事件转发
*/
func (o factoryObserve) start() {
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
func (o factoryObserve) Register(msgId int, f factory.IFactory) {
	o.registerChan <- FactoryRegister{MsgId: msgId, Factory: f}
}

/*
消息注册内部接口,去除锁
*/
func (o factoryObserve) register(r FactoryRegister) {
	o.registerCenter[r.MsgId] = r.Factory
}

/*
GetTasker 获取消息外部接口
*/
func (o factoryObserve) GetTasker(msgId int) task.Tasker {
	o.getTaskerChan <- msgId
	return <-o.sendTaskerChan
}

/*
获取tasker,通过factory生产tasker,根据msgId策略执行任务(因此这个模块用到了观察者、工厂、策略模式)
*/
func (o factoryObserve) sendTasker(msgId int) {
	f, exist := o.registerCenter[msgId]
	if !exist {
		o.sendTaskerChan <- nil
		return
	}
	o.sendTaskerChan <- f.CreateFactory()
}

/*
Exit 退出
*/
func (o factoryObserve) Exit() {
	go func() {
		o.exit <- struct{}{}
	}()
}
