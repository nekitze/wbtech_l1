package main

import (
	"fmt"
	"strings"
)

func IsAllSymbolsUnique(s string) bool {
	s = strings.ToLower(s)
	runes := make(map[rune]bool)
	for _, r := range s {
		if runes[r] {
			return false
		}
		runes[r] = true
	}

	return true
}

func main() {
	s := []string{"abcd", "abCdefAaf", "aabcd", "Test123", "123345"}

	for _, str := range s {
		fmt.Println(IsAllSymbolsUnique(str))
	}
}
