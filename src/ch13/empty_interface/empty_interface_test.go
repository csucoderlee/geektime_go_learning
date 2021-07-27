package empty_interface_test

import (
	"fmt"
	"testing"
)

func doSomething(p interface{}) {
	switch v := p.(type) {

	case int:
		fmt.Println("Integer", v)

	case string:
		fmt.Println("String", v)

	default:
		fmt.Println("Unknown Type")
	
	}
}

func TestDoSomething(t *testing.T) {
	doSomething(10)
}
