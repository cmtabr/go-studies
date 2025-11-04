package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, char := range s {
		if lastIndex, found := charIndexMap[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndexMap[char] = i
		fmt.Println((charIndexMap))
		if currentLength := i - start + 1; currentLength > maxLength {
			fmt.Println((currentLength))
			maxLength = currentLength
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
