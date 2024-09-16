package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 && (x%10 == 0 && x != 0) {
		return false
	}

	s := strconv.Itoa(x)
	var r string

	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}

	return s == r
}

func isPalindromeOptimized(x int) bool {
	if x < 0 && (x%10 == 0 && x != 0) {
		return false
	}

	var r int
	for x > r {
		r = r*10 + x%10
		x /= 10

		fmt.Println(r)
		fmt.Println(x)
	}

	return x == r || x == r/10
}

func main() {
	fmt.Println(isPalindromeOptimized(121))
}
