package main

import "fmt"

func swap(x *int, y *int) {
	var k int

	k = *x
	*x = *y
	*y = k
}

func main() {
	a := 10
	b := &a

	fmt.Println(a, b)

	// check the data type of pointer
	fmt.Printf("%T\n", b)

	// read value using pointer
	fmt.Println(*b)

	// swapping of two numbers using pointers
	x := 10
	y := 20

	fmt.Printf("before swap x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("after swap x = %d, y = %d", x, y)
}
