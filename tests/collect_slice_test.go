package tests

import (
	"github.com/alicfeng/gutilib/collection"
	"testing"
)

func TestCollectSliceFilter(t *testing.T) {
	data := []int64{1, 2, 3}

	result := collection.SliceFilter(data, func(index int, item int64) bool {
		return item == 2
	})

	if 1 != len(result) {
		t.Fail()
	}
}

func TestCollectSliceSliceSub(t *testing.T) {
	data := []int{1, 2, 3}

	result := collection.SliceSub(data, 0, 1)

	if 2 != len(result) {
		t.Fail()
	}
}
