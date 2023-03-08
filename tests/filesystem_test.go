package tests

import (
	"github.com/alicfeng/gutilib/filesystem"
	"runtime"
	"testing"
)

func TestPathExist(t *testing.T) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return
	}

	if !filesystem.PathExist(file) {
		t.Fail()
	}
}
