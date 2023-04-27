package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	//使用装饰器来动态地给组件添加新的行为
	// 创建一个组件类，并对该类添加装饰器
	component := &ConcreteComponent{}
	decoratorA := &ConcreteDecoratorA{component: component}
	fmt.Println(decoratorA.Operation())
}
