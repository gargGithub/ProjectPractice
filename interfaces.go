package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}
type circle struct {
	rad float64
}

func(s square) area() float64{
	return s.side*s.side
}

func(c circle) area() float64{
	return math.Pi*c.rad*c.rad
}

type Shapes interface {
	area() float64
}

func totalArea(shape Shapes){
	fmt.Println(shape)
	fmt.Println(shape.area())

}
func main() {


	p1:=square{10}
	totalArea(p1)
	p2:=circle{10}
	totalArea(p2)
}
