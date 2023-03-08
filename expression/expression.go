package expression

// Ternary 基于泛型实现三元表达式
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}

	return falseValue
}
