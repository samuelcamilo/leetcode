package main

import "fmt"

func sortedSquares(nums []int) []int {
	arr := make([]int, len(nums))

	firstIdx := 0
	lastIdx := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		firstItem := nums[firstIdx] * nums[firstIdx]
		lastItem := nums[lastIdx] * nums[lastIdx]

		if lastItem >= firstItem {
			arr[i] = lastItem
			lastIdx--
		} else {
			arr[i] = firstItem
			firstIdx++
		}
	}

	return arr
}

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(arr))
}
