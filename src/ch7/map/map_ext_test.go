package map_ext_test

import "testing"

//map的value可以是方法
func TestMapWithFuc(t *testing.T) {
	map1 := map[string]func(n int) int{}
	map1["one"] = func(op int) int { return op }
	map1["two"] = func(op int) int { return op * op }

	t.Log(map1["one"](1), map1["two"](2))

}

//map实现set
func TestMap2Set(t *testing.T) {
	map1 := map[string]bool{}
	map1["one"] = true

	if map1["one"] {
		t.Log(" one is already in set")
	} else {
		t.Log(" one is not in set")
	}

	delete(map1, "one")

	if map1["one"] {
		t.Log(" one is already in set")
	} else {
		t.Log(" one is not in set")
	}

}
