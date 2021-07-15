package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := true; a {
		t.Log(a)
	}
}

func TestSwitchMultiCase(t *testing.T) {

	for n := 0; n < 5; n++ {
		switch n {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for n := 0; n < 5; n++ {
		switch {
		case n%2 == 0:
			t.Log("Even")
		case n%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unkonwn")

		}
	}
}
