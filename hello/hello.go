package main

import (
	"fmt"

	"github.com/aaqib/golang_basics/greetings"
)

func main() {
	message := greetings.Hello("aaqib")

	fmt.Println(message)
}
