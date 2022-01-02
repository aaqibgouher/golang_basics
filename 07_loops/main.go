package main

import "fmt"

func main() {
	// long method
	i := 1

	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// shortcut method
	// for i = 1; i <= 100; i++ {
	// 	fmt.Println(i)
	// }

	// FizzBuzz
	for i = 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
