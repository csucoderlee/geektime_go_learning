package type_test

import (
	"log"
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64 = 1

	b = int64(a)

	var c MyInt = 64
	c = MyInt(b)

	t.Log(a, b, c)

	t.Log(math.MaxFloat32, math.MaxInt32, math.MaxInt64)

}

func TestPoint(t *testing.T) {
	a := 1
	b := &a

	t.Log(a, b)

	t.Logf("%T %T", a, b)
	log.Printf("%T %T", a, b)
}
