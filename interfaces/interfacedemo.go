package interfaces

import (
	"fmt"
	"math"
)

// icinde iki adet metod kalıbı
// barındıran geometry isimli bir
// interface tanimladik

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// reqtangle methods
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area : ", g.area())
	fmt.Println("perim: ", g.perim())
}

func IntTest() {
	r := rect{width: 7, height: 11}
	c := circle{radius: 13}

	fmt.Println("rectangle : ")
	measure(r)

	fmt.Println("circle :")
	measure(c)
}
