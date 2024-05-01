package main

import (
	"fmt"
)

// Brute Force...
func twoSumBruteForce(nums []int, target int) []int {
	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSumOptimized(nums []int, target int) []int {
	hm := make(map[int]int)

	for i, num := range nums {
		_, ok := hm[num]
		if ok {
			return []int{hm[num], i}
		}

		hm[target-num] = i
	}

	return nil
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 26

	output := twoSumOptimized(input, target)

	fmt.Println(output)
}
