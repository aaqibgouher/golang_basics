package main

import "fmt"

func main() {
	// var fruits [3]string
	// fruits[0] = "mango"
	// fruits[1] = "banana"
	// fruits[2] = "apple"
	fruits := []string{"mango", "banana", "apple", "pine apple"}

	fmt.Println(fruits)
	fmt.Println(len(fruits)) /*get the length of an array*/
	fmt.Println(fruits[0:2]) /*slicling array element starting from index 0 and till 2(excluded)*/
}
