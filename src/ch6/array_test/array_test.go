package array_test

import "testing"

//数组的初始化
func TestArrayInit(t *testing.T) {
	var a [3]int
	t.Log(a[1], a[2])

	array1 := [3]int{1, 2, 3}
	t.Log(array1[0], array1[1], array1[2])

	array2 := [...]int{1, 2, 3, 4}
	t.Log(array2[0], array2[1], array2[2])
}

//数组的遍历
func TestArrayTravel(t *testing.T) {
	array2 := [...]int{1, 2, 3, 4}

	for i := 0; i < len(array2); i++ {
		t.Log(array2[i])
	}

	for index, num := range array2 {
		t.Log(index, num)
	}

	for _, num := range array2 {
		t.Log(num)
	}

	for index, _ := range array2 {
		t.Log(index)
	}
}

//数组的截取  a[开始索引(包含) : 结束索引(不包含)]
func TestArraySection(t *testing.T) {

	array1 := [...]int{1, 2, 3, 4}

	array2 := array1[1:]
	t.Log(array2)

}

// 数组比较
func TestArrayCompare(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	// 长度不同，不能比较
	//t.Log(a == c)
	t.Log(a == d)
}
