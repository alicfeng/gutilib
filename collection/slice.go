package collection

// SliceFilter 提供闭包函数作过滤参数.
func SliceFilter[S any](in []S, predicate func(index int, item S) bool) []S {
	result := make([]S, 0, len(in))

	for index, item := range in {
		if false == predicate(index, item) {
			continue
		}
		result = append(result, item)
	}

	return result
}

// SliceSub 切片截取 start 开始位置索引 end 结束位置索引
func SliceSub[S any](in []S, start, end int) []S {
	result := make([]S, len(in))

	copy(result, in)

	result = result[start : end+1 : len(in)]

	return result
}
