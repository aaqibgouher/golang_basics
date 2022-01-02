package main

import "fmt"

func sayHello(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(sayHello("Aaqib"))
	fmt.Println(getSum(10, 12))
}
