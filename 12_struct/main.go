package main

import "fmt"

type User struct {
	name   string
	email  string
	maths  int
	cs     int
	stats  int
	status string
}

// struct method (value receiver)
func (user User) getStatus() string {
	return user.status
}

// struct method (pointer receiver)
func (user *User) getReasult() map[string]string {
	if user.maths < 33 && user.cs < 33 && user.stats < 33 {
		user.status = "fail"
	} else {
		user.status = "pass"
	}

	return map[string]string{
		"name":   user.name,
		"email":  user.email,
		"status": user.status,
		"result": "your result is out!",
	}
}

func main() {
	user := User{
		name:   "Aaqib",
		email:  "aaqibgouher@gmail.com",
		maths:  34,
		cs:     33,
		stats:  32,
		status: "false",
	}

	fmt.Println(user)
	fmt.Println(user.getStatus())
	data := user.getReasult()
	fmt.Println(data)
}
