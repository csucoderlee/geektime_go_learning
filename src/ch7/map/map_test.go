package map_ext_test

import "testing"

func TestMapInit(t *testing.T) {

	map1 := map[string]int{"one": 1, "two": 2}
	t.Log(map1, len(map1))

	var map2 = map[string]int{}
	map2["one"] = 1
	t.Log(map2, len(map2))

	map3 := make(map[string]int, 10)
	t.Log(map3, len(map3))
}

func TestMapIsNotExistKey(t *testing.T) {

	var map1 = map[string]int{}
	//不存在的key，获取会返回0值
	t.Log(map1["one"])

	if v, ok := map1["one"]; ok {
		t.Log("key one exist", v)
	} else {
		t.Log("key one not exist")
	}

	map1["two"] = 0

	if v, ok := map1["two"]; ok {
		t.Log("key two exist, ths value is ", v)
	} else {
		t.Log("key two not exist")
	}
}

func TestMapTravel(t *testing.T) {
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}

	for k, v := range map1 {
		t.Log(k, v)
	}
}
