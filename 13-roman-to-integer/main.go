package main

import "fmt"

var hm = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// examples: MCMXCIV
func romanToInt(s string) int {
	sum := 0
	bigest := 0
	for i := len(s) - 1; i >= 0; i-- {
		value := hm[rune(s[i])]
		if value >= bigest {
			sum = sum + value
			bigest = value
			continue
		}

		sum = sum - value
	}

	return sum
}

func main() {
	result := romanToInt("MCMXCIV")
	fmt.Println(result)
}
