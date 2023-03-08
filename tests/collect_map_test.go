package tests

import (
	"github.com/alicfeng/gutilib/collection"
	"testing"
)

func mockMapData() map[string]int64 {
	return map[string]int64{
		"math":    88,
		"english": 90,
		"chinese": 60,
	}
}

func TestCollectionMapFilter(t *testing.T) {
	result := collection.MapFilter(mockMapData(), func(key string, value int64) bool {
		return value > 80
	})
	if 2 != len(result) {
		t.Fail()
	}
}

func TestCollectionMapKeys(t *testing.T) {
	result := collection.MapKeys(mockMapData())
	if 3 != len(result) {
		t.Fail()
	}
}

func TestCollectionMapValues(t *testing.T) {
	result := collection.MapKeys(mockMapData())
	if 3 != len(result) {
		t.Fail()
	}
}

func TestCollectionMapIsContain(t *testing.T) {
	if collection.MapIsContain(mockMapData(), 99) {
		t.Fail()
	}
}

func TestCollectionMapExistKey(t *testing.T) {
	if false == collection.MapExistKey(mockMapData(), "math") {
		t.Fail()
	}
}
