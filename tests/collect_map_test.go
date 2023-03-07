package tests

import (
	"github.com/alicfeng/gutilib/src/collection"
	"testing"
)

func TestCollectionMapFilter(t *testing.T) {
	data := map[string]int64{
		"math":    88,
		"english": 90,
		"chinese": 60,
	}

	result := collection.MapFilter(data, func(key string, value int64) bool {
		return value > 80
	})
	if 2 != len(result) {
		t.Fail()
	}
}

func TestCollectionMapKeys(t *testing.T) {
	data := map[string]int64{
		"math":    88,
		"english": 90,
		"chinese": 60,
	}

	result := collection.MapKeys(data)
	if 3 != len(result) {
		t.Fail()
	}
}

func TestCollectionMapIsContain(t *testing.T) {
	data := map[string]int64{
		"math":    88,
		"english": 90,
		"chinese": 60,
	}

	if collection.MapIsContain(data, 99) {
		t.Fail()
	}
}

func TestCollectionMapExistKey(t *testing.T) {
	data := map[string]int64{
		"math":    88,
		"english": 90,
		"chinese": 60,
	}

	if false == collection.MapExistKey(data, "math") {
		t.Fail()
	}
}
