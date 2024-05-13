package main

import "fmt"

func createSet(strs []string) map[string]bool {
	set := make(map[string]bool)
	for _, str := range strs {
		set[str] = true
	}
	return set
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	strsSet := createSet(strs)

	for key := range strsSet {
		fmt.Println(key)
	}
}
