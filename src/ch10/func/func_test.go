package func_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFunc(t *testing.T) {

	result1, result2 := multiResult()
	t.Log(result1, result2)

	result3 := paramFuc(func(op int) int {
		return op * 3
	})
	t.Log(result3)

	result4 := paramFuc(func(op int) int {
		return op * 4
	})
	t.Log(result4)

	result5 := paramFuc(funcAsParam)
	t.Log(result5)
}

//函数可以有多个返回值
func multiResult() (int, int) {
	return rand.Int(), rand.Int()
}

func paramFuc(inner func(op int) int) int {
	return inner(1)
}

// funcAsParam函数可以作为 paramFuc的参数， 见上方result5
func funcAsParam(op int) int {
	return op * 5
}

func TestSomeParam(t *testing.T) {
	t.Log(someParam(1, 2, 3, 4))
	t.Log(someParam(1, 2, 3))

}
func someParam(op ...int) int {
	result := 0
	for _, value := range op {
		result += value
	}
	return result
}

func clear() {
	fmt.Println("clear resources")
}

func TestDefer(t *testing.T) {
	defer clear()
	t.Log("start before clear")

	//defer 仍会执行
	panic("error")
}
