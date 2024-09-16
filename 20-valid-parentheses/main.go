package main

import "fmt"

func isValid(str string) bool {
	st := []rune{}
	hm := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}

	index := 0
	for _, s := range str {
		if _, ok := hm[s]; ok {
			st = append(st, s)
			index++
			continue
		}

		fmt.Printf("%v", st[3])

		if len(st) > 0 && s == hm[st[index-1]] {
			st = st[:index-1]
			fmt.Printf("%v\n", st)
			index--
		} else {
			return false
		}
	}

	return len(st) == 0
}

func main() {
	input := "({[]})"
	output := isValid(input)

	fmt.Println(output)
}

// debug...

// 1. st[0] = '('
// 2. st[1] = '{'
// 3. st[2] = '['
