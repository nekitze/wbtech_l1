package main

import (
	"fmt"
	"runtime"
	"strings"
)

var justString string

func createHugeString(length int) string {
	return fmt.Sprintf("%s", make([]byte, length))
}

func someFunc() {
	v := createHugeString(1 << 20)
	/*
		При создании среза сохраняется ссылка на большую
		исходную строку, поэтому она не будет собрана GC.
	*/
	//justString = v[:100]

	// Лучше так
	justString = strings.Clone(v[:100])
}

func main() {
	var m runtime.MemStats
	someFunc()

	runtime.GC()

	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v KiB", bToKb(m.Alloc))
	fmt.Printf("\nTotalAlloc = %v KiB", bToKb(m.TotalAlloc))
	fmt.Printf("\nNumGC = %v\n", m.NumGC)
}

func bToKb(b uint64) uint64 {
	return b / 1024
}
