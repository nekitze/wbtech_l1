package main

import (
	"fmt"
)

func reverse(s string) string {
	b := []rune(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func main() {
	s1 := "главрыба"
	s2 := "෴இ߷"
	s3 := "½ ⅓ ⅔ ¼ ¾ ⅕ ⅖ ⅗ ⅘ ⅙ ⅚ ⅐ ⅛ ⅜ ⅝ ⅞ ⅑ ⅒"

	s1 = reverse(s1)
	s2 = reverse(s2)
	s3 = reverse(s3)

	fmt.Printf("%s\n%s\n%s\n", s1, s2, s3)
}
