package src

type Publish struct {
	name     string
	Follower []*Subscriber //存放所有订阅者
	Msg      interface{}
}

// 创建一个发布者
func NewPub(name string) *Publish {
	return &Publish{
		name:     name,
		Follower: make([]*Subscriber, 0),
	}
}

// 添加消息
func (p *Publish) Add(msg interface{}) {
	p.Msg = msg
	// 将消息通知给所有订阅者
	p.Notify()
}

func (p *Publish) Notify() {
	for _, v := range p.Follower {
		v.Subs[p] <- p.Msg
	}
}

// 添加订阅者到切片中
func (p *Publish) AddSub(s *Subscriber) {
	p.Follower = append(p.Follower, s)
}

func (p *Publish) UnSub(s *Subscriber) {
	// 从切片中移除订阅者
	for i, v := range p.Follower {
		if v == s {
			p.Follower = append(p.Follower[:i], p.Follower[i+1])
		}
	}
}
