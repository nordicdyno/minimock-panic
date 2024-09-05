package minipanic

import (
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
)

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	mc := minimock.NewController(t)
	getterMock := NewGetterMock(mc)

	getterMock.GetMock.Return(fmt.Errorf("fail"))
	getterMock.Get2Mock.Return()
	gc := GetterCaller{getterMock}

	gc.Run()

	t.Log("test finalized")
}
