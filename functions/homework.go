package main

func Map(data []int, action func(int) int) []int {
	if len(data) == 0 {
		return data
	}

	copyData := make([]int, 0, len(data))
	for _, val := range data {
		copyData = append(copyData, action(val))
	}

	return copyData
}

func Filter(data []int, action func(int) bool) []int {
	if len(data) == 0 {
		return data
	}

	copyData := make([]int, 0, len(data))
	for _, val := range data {
		if action(val) {
			copyData = append(copyData, val)
		}
	}

	return copyData
}

func Reduce(data []int, initial int, action func(int, int) int) int {
	if len(data) == 0 {
		return 0
	}

	var impl func(initial int) int
	impl = func(n int) int {
		for _, val := range data {
			n = action(n, val)
		}

		return n
	}

	return impl(initial)
}
