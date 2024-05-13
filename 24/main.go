package main

import (
	"fmt"
	"wbtech_l1/24/point"
)

func main() {
	p1 := point.NewPoint(-1, 3)
	p2 := point.NewPoint(6, -2)

	fmt.Println(p1.GetDistance(p2))
}
