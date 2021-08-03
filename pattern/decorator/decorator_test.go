package decorator

import "testing"

/*
装饰器类
*/
func TestDecorator(t *testing.T) {
	g := Gateway{}
	if err := TryDecorator(g, "test", 3); err != nil {
		t.Fatal(err)
	}
	t.Fatal("")
	return
}

/*
装饰器函数
*/
func TestDecoratorFunc(t *testing.T) {
	if err := TryFunc(sendGateway, "test", 3); err != nil {
		t.Fatal(err)
	}
	if err := TryFunc(addData, "test", 5); err != nil {
		t.Fatal(err)
	}
	t.Fatal("")
	return
}
