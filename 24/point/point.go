package point

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) GetDistance(o *Point) float64 {
	x := math.Pow(o.x-p.x, 2)
	y := math.Pow(o.y-p.y, 2)
	return math.Sqrt(x + y)
}
