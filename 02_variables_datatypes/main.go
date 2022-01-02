package main

import "fmt"

func main() {
	// 1. using var
	// var name string = "Aaqib"		/*both way we can define*/
	// var name = "Aaqib"
	name := "Aaqib" /*shorthand*/
	weight := 1.4

	var age = 20
	var status = true
	status = false

	// constants declaration
	const isDone = true

	city, email := "bangalore", "aaqib@gmail.com"

	fmt.Println(name, age, status, isDone, weight, city, email)
	// to get the datatype of variale, we will use Print
	fmt.Printf("name is %T, age is %T, status is %T, wright is %T", name, age, status, weight)
}
