package main

import (
	f "PubAndSub-pattern/src"
	"time"
)

func main() {
	p1 := f.NewPub("阿松大")
	s1 := f.NewSub("Lisa")
	s2 := f.NewSub("chd")
	s3 := f.NewSub("tony")

	s1.Subscribe(p1)
	s2.Subscribe(p1)

	p1.Add("test1")

	time.Sleep(3 * time.Second)

	s1.UnSubscribe(p1)
	s3.Subscribe(p1)

	time.Sleep(2 * time.Second)
	p1.Add("test1")

	for {

	}
}
