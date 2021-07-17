package slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var slice1 []int
	t.Log(len(slice1), cap(slice1))

	slice1 = append(slice1, 1)
	t.Log(len(slice1), cap(slice1))

	slice2 := []int{1, 2, 3, 4, 5}
	t.Log(len(slice2), cap(slice2))

	slice3 := make([]int, 3, 5)
	t.Log(slice3[0], slice3[1], slice3[2], len(slice3), cap(slice3))
	slice3 = append(slice3, 1)
	t.Log(slice3[0], slice3[1], slice3[2], slice3[3], len(slice3), cap(slice3))

}

func TestSliceGrowing(t *testing.T) {
	var slice1 = []int{}
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
		t.Log(len(slice1), cap(slice1))
	}
}

//切片共享存储结构
func TestSliceShareMemory(t *testing.T) {
	month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	Q2 := month[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	Q2[0] = "Unknown"
	t.Log(month[3])
}
