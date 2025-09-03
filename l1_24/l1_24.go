package l1_24

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (cp *Point) Distance(other Point) float64 {
	xDistance := math.Abs(cp.x - other.x)
	yDistance := math.Abs(cp.y - other.y)
	return math.Sqrt(xDistance*xDistance + yDistance*yDistance)
}

func MainFunc() {
	point1 := NewPoint(10, 20)
	point2 := NewPoint(20, 30)
	fmt.Println(point1.Distance(*point2))
}
