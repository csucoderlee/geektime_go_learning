package string_test

import "testing"

func TestStr(t *testing.T) {

	//byte slice
	s := "中国"

	//只读的slice，不可变
	//s[1] = "华"

	s = "中华"

	//字节数，slice的length，即为byte的大小
	t.Log(len(s))

	//字符数
	t.Log(len([]rune(s)))
	t.Log(len([]int32(s)))
}
