package main

import (
	"fmt"
	"math"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func main() {
	// c := Circle{23, 45, 66}
	// fmt.Println(circleArea(&c))
	// fmt.Println(c.area())
	// var cs []Circle
	// cs = append(cs, Circle{23, 34, 45})
	// cs = append(cs, Circle{65, 32, 21})

	// var rs []Rectangle
	// rs = append(rs, Rectangle{23, 34, 45, 99})
	// rs = append(rs, Rectangle{65, 32, 21, 0})

	fmt.Println(totalArea(&Circle{23, 34, 45}, &Rectangle{65, 32, 21, 0}))

	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{33, 45, 66},
			&Rectangle{77, 88, 99, 33},
		},
	}
	fmt.Println(totalArea(&multiShape))

}

// ----------------------------------------------------------- //
// ----------------------------------------------------------- //
// func circleArea(c *Circle) float64 {
// 	return math.Pi * c.r * c.r
// }

// func totalArea(ciscles ...Circle) float64 {
// 	var total float64
// 	for _, c := range ciscles {
// 		total += c.area()
// 	}
// 	return total
// }

// func rectangleArea(x1, x2, y1, y2 float64) float64 {
// 	l := distance(x1, y1, x1, y2)
// 	w := distance(x1, y1, x2, y1)
// 	return l * w
// }
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// func totalArea(ciscles []Circle, rectangles []Rectangle) float64 {
// 	var total float64
// 	for _, c := range ciscles {
// 		total += c.area()
// 	}
// 	for _, r := range rectangles {
// 		total += r.area()
// 	}
// 	return total
// }

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
