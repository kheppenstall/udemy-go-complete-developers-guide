package shape

import "fmt"

type shape interface {
	getArea() float64
}

type rectangle struct {
	height float64
	width  float64
}

type triangle struct {
	height float64
	width  float64
}

func (r rectangle) getArea() float64 {
	return r.height * r.width
}

func (t triangle) getArea() float64 {
	return t.height * t.width / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
