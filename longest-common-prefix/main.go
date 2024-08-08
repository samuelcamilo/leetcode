package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	minor := len(strs[0])

	for _, s := range strs[1:] {
		if len(s) <= minor {
			minor = len(s)
		}
	}

	occr := 0
	common := strs[0][:minor]
	times := minor
	i := 0

	for i < times {
		for _, str := range strs {
			if common != str[:minor] {
				occr = 0
				break
			}
			occr++
		}

		if occr == len(strs) {
			return common
		}

		minor = minor - 1
		common = common[:minor]
		i = i + 1
	}

	return common
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
