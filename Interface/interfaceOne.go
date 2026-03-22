package Interface

import (
	"fmt"
	"math"
)

type Figure interface {
	Area() float64
	Perimetr() float64
}

type Cirkle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Box struct {
	Len float64
}

func (c *Cirkle) Perimetr() float64 {
	fmt.Print("Периметр круга: ")
	return 2 * math.Pi * c.Radius
}

func (c *Rectangle) Perimetr() float64 {
	fmt.Print("Периметр прямоугольника: ")
	return 2 * (c.Height + c.Width)
}

func (c *Box) Perimetr() float64 {
	fmt.Print("Периметр квадрата: ")
	return 4 * c.Len
}

func (c *Cirkle) Area() float64 {
	fmt.Print("Площадь круга: ")
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Rectangle) Area() float64 {
	fmt.Print("Площадь прямоугольника: ")
	return c.Height * c.Width
}

func (c *Box) Area() float64 {
	fmt.Print("Площадь квадрата: ")
	return math.Pow(c.Len, 2)
}

func Use(f Figure) {
	fmt.Println(f.Perimetr())
	fmt.Println(f.Area())
}
