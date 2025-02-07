package main

import "fmt"

type User struct {
	name     string
	age      int
	email    string
	location struct {
		country string
	}
}

func create_user(name string, age int, email string) (new_user User) {
	new_user.name = name
	new_user.age = age
	new_user.email = email
	new_user.location.country = "Pakistan"
	return new_user
}

func set_info(user *User) {
	fmt.Print("Enter your name: ")
	fmt.Scan(&user.name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&user.age)
	fmt.Printf("Enter your email address: ")
	fmt.Scan(&user.email)
}

func get_info(user User) {
	details := fmt.Sprintf("%q is %d years old who using %q as email address", user.name, user.age, user.email)
	fmt.Println(details)
}

func main() {
	fmt.Println("Struct in GO")
	var hassan User
	set_info(&hassan)
	get_info(hassan)
	test := create_user("hassan", 20, "hassan@test.com")
	ali := create_user("ali", 10, "ali@gmail.com")
	get_info(test)
	get_info(ali)
}
