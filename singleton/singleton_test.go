package singleton

import (
	"fmt"
	"testing"
)

func TestLazy(t *testing.T) {
	lazy1 := getLazyInstance()
	lazy2 := getLazyInstance()
	fmt.Println(lazy1 == lazy2) //true
}

func TestHungry(t *testing.T) {
	hungry1 := getHungryInstance()
	hungry2 := getHungryInstance()
	fmt.Println(hungry1 == hungry2) //true
}
