package main

import "fmt"

func main() {
	cars := []string{"bmw", "audi", "range rower"}

	// with index
	for index, car := range cars {
		fmt.Printf("%d => %s\n", index, car)
	}

	// without index
	for _, car := range cars {
		fmt.Println(car)
	}

	// range with map
	user := map[string]string{
		"name":  "aaqib",
		"email": "aaqib@gmail.com",
		"age":   "20",
	}

	for key, value := range user {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}

}
