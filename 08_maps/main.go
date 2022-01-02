package main

import "fmt"

func main() {
	// define a map
	user := make(map[string]string)

	// assigning
	user["name"] = "Aaqib Gouher"
	user["email"] = "aaqibgouher@gmail.com"
	user["age"] = "20"

	fmt.Println(user, user["name"])

	// map with integer key
	data := make(map[int]string)
	data[0] = "Aaqib"
	data[1] = "Gouher"

	fmt.Println(data, len(data))

	// deleting a key from map
	delete(user, "age")
	fmt.Println(user)

	// declare and initialize the map together
	person := map[string]string{
		"name":  "Aaqib",
		"email": "aaqibgouher@gmail.com",
	}

	fmt.Println(person)
}
