package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	slices.Reverse(ss)

	builder := strings.Builder{}
	for _, str := range ss {
		builder.WriteString(str)
		builder.WriteRune(' ')
	}

	return builder.String()
}

func main() {
	s := "snow dog sun"
	s = reverseWords(s)
	fmt.Println(s)
}
