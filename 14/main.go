package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := []interface{}{1, "stroka", true, make(chan int)}

	for _, d := range data {
		// use reflect package to determine type
		t := reflect.TypeOf(d)
		fmt.Println(t)
	}
}
