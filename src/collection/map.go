package collection

// MapFilter 提供闭包函数作过滤参数.
func MapFilter[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) map[K]V {
	result := map[K]V{}

	for key, value := range in {
		if false == predicate(key, value) {
			continue
		}
		result[key] = value
	}

	return result
}

// MapExistKey 断言是否存在指定的键.
func MapExistKey[K comparable](in map[K]interface{}, key K) bool {
	_, ok := in[key]

	return ok
}

// MapKeys 提取所有的键名.
func MapKeys[K comparable](in map[K]interface{}) []K {
	keys := make([]K, len(in))

	for key := range in {
		keys = append(keys, key)
	}

	return keys
}

// MapIsContain 断言指定值是否被包含
func MapIsContain[V comparable](in map[interface{}]V, predicate V) bool {
	for _, value := range in {
		if value == predicate {
			return true
		}
	}

	return false
}
