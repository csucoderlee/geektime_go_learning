package fib

import (
	"testing"
)

func TestFib(t *testing.T) {

	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 6; i++ {
		t.Log(b)
		temp := a
		a = b
		b = temp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
