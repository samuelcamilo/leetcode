package main

import (
	"fmt"
)

// Brute Force
func twoSum(nums []int, target int) []int {
	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 26

	output := twoSum(input, target)

	fmt.Println(output)
}
