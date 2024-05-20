package main

import "fmt"

func setBit(num, val, pos int64) int64 {
	if val == 1 {
		// 1 << pos - битовый сдвиг единицы на pos лево
		// | - битовое ИЛИ
		return num | 1<<pos
	} else {
		// &^ битовое исключаюеще ИЛИ
		return num &^ (1 << pos)
	}
}

func main() {
	var a int64 = 111
	fmt.Printf("%b\n", a) // 1101111

	fmt.Printf("%b\n", setBit(a, 0, 0)) // 1101110
	fmt.Printf("%b\n", setBit(a, 0, 1)) // 1101101
	fmt.Printf("%b\n", setBit(a, 0, 2)) // 1101011
	fmt.Printf("%b\n", setBit(a, 0, 3)) // 1100111
	fmt.Printf("%b\n", setBit(a, 0, 4)) // 1101111
	fmt.Printf("%b\n", setBit(a, 0, 5)) // 1001111
	fmt.Printf("%b\n", setBit(a, 0, 6)) // 101111

	fmt.Printf("%b\n", setBit(a, 1, 4)) // 1111111
}
