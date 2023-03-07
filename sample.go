package main

import (
	"github.com/alicfeng/gutilib/src/expression"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		at := time.Now().Second()
		va := expression.Ternary(at%2 == 0, "偶数", "奇数")
		println(va)
	}
}
