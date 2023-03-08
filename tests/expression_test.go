package tests_test

import (
	"github.com/alicfeng/gutilib/expression"
	"testing"
	_ "testing"
)

func TestTernary(t *testing.T) {

	if 1 != expression.Ternary(true, 1, 0) {
		t.Fail()
	}

	if 0 != expression.Ternary(false, 1, 0) {
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
