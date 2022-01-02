package main

import "fmt"

func main() {
	color := "green"

	// common practise is not use () while writing the conditions
	// if color == "red" {
	// 	fmt.Println("you have selected red")
	// } else if color == "green" {
	// 	fmt.Println("you have selected green")
	// } else {
	// 	fmt.Println("you have not selected any color")
	// }

	// switch
	switch color {
	case "red":
		fmt.Println("you have selected red")
	case "green":
		fmt.Println("you have selected green")
	default:
		fmt.Println("you have not selected any color")
	}
}
