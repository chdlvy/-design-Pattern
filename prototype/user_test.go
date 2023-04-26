package prototype

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	u1 := User{
		Name: "ccc",
		Age:  18,
		M: map[string]string{
			"hhh": "32121",
		},
	}
	u2 := u1.Clone()
	fmt.Println(u2)
	u2.M["hhh"] = "qwe"

	fmt.Println("----------------")
	fmt.Println(u1)
	fmt.Println(u2)
}
