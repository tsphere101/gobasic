package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello")
	fmt.Println(isPalindrome(-303))

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i < (len(s)/2)+1; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
