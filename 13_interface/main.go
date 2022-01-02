package main

import "fmt"

type Shape interface {
	area() int
}

type Square struct {
	side int
}

type Rectangle struct {
	length, breadth int
}

func (sq Square) area() int {
	return sq.side * sq.side
}

func (rect Rectangle) area() int {
	return rect.length * rect.breadth
}

func getArea(s Shape) int {
	return s.area()
}

func main() {
	sq := Square{
		side: 4,
	}

	rect := Rectangle{
		length:  10,
		breadth: 2,
	}

	fmt.Printf("area of sqaure is %d\n", getArea(sq))
	fmt.Printf("area of rectangle is %d\n", getArea(rect))
}
