package decorator

import "fmt"

type Component interface {
	Operation() string
}

//组件类
type ConcreteComponent struct{}

//实现Component接口
func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// 结构体实现Component接口
//装饰器
type ConcreteDecoratorA struct {
	component Component //被装饰的组件
}

func (d *ConcreteDecoratorA) Operation() string {
	//调用被装饰的组件的方法
	d.component.Operation()
	//do something.......
	return fmt.Sprintf("我是ConcreteDecoratorA，是组件%s的装饰器", d.component.Operation())
}

// type ConcreteDecoratorB struct {
// 	component Component
// }

// func (d *ConcreteDecoratorB) Operation() string {
// 	return fmt.Sprintf("我是ConcreteDecoratorB，是组件%s的装饰器", d.component.Operation())
// }
