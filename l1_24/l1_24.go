package l1_24

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func Distance(p, q Point) float64 {
	xDistance := math.Abs(q.x - p.x)
	yDistance := math.Abs(q.y - p.y)
	return math.Sqrt(xDistance*xDistance + yDistance*yDistance)
}

func MainFunc() {
	point1 := Point{x: 10, y: 20}
	point2 := Point{x: 5, y: 10}
	fmt.Println(Distance(point1, point2))
}
