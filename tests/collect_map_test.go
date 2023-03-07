package tests

import (
	"fmt"
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

	fmt.Print(result)
}
