package main

import "fmt"

func twoSum(nums []int, target int) []int {
	arr := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			res := nums[i] + nums[j]
			if res == target && i != j {
				arr = append(arr, i)
				arr = append(arr, j)

				return arr
			}
		}
	}

	return arr
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 26

	output := twoSum(input, target)

	fmt.Println(output)
}
