package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return 2 * (r.height + r.width)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}