package src

import "fmt"

type Subscriber struct {
	name string
	Subs map[*Publish]chan interface{} //保存所有订阅了的发布者，采用map是因为经常要读操作，用map更快
}

//创建订阅者
func NewSub(name string) *Subscriber {
	return &Subscriber{
		name: name,
		Subs: make(map[*Publish]chan interface{}),
	}
}

//订阅操作
func (s *Subscriber) Subscribe(p *Publish) {
	//将订阅者添加到Subs中
	c := make(chan interface{})
	s.Subs[p] = c
	p.AddSub(s)
	//开启协程一直监听发布者信息
	go Listen(s, p)
}

func Listen(s *Subscriber, p *Publish) {
	//不断监听channel，接收订阅者信息
	for {
		if msg, ok := <-s.Subs[p]; ok {
			fmt.Printf("这里是%s，收到来自%s发送的：%s\n", s.name, p.name, msg.(string))
		} else {
			fmt.Println("退订")
			break
		}

	}
}

func (s *Subscriber) UnSubscribe(p *Publish) {
	//关闭channel
	close(s.Subs[p])
	p.UnSub(s)
	// 删除键值对
	delete(s.Subs, p)
}
