package main

import "fmt"

func main() {

	fmt.Println("starting..")
	arr := []int{1, 1, 3, 4, 5}

	fmt.Println("result ", removeDuplicates(arr))

}

func removeDuplicates(nums []int) []int {

	var alreadyHas []int
	var count int = 0

	for i := range nums {
		num := nums[i]
		if !HasNum(alreadyHas, num) {
			alreadyHas = append(alreadyHas, num)
		} else {
			count++
		}

	}
	return alreadyHas
}

func HasNum(arr []int, num int) bool {
	for index := range arr {
		if arr[index] == num {
			return true
		}
	}
	return false
}
