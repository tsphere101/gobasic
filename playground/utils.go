package main

func reduce(nums []int, method func(int, int) int) int {

	ag := 0
	// iteration
	for _, element := range nums {
		ag = method(ag, element)
	}
	return ag
}
