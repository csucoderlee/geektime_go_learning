package panic_recover_test

import (
	"github.com/pkg/errors"
	"testing"
)

func TestPanicRecover(t *testing.T) {

	defer func() {

		err := recover()

		if err != nil {
			t.Log("panic")
		}
	}()

	panic(errors.New("exit"))
}
