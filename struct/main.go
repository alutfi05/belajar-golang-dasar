package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	isActive  bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Ahmad Lutfi"
	user.LastName = "Rizki Patria"
	user.Email = "aluthfi0522@gmail.com"
	user.isActive = true

	user2 := User{
		ID: 2,
		FirstName: "Laila Indah", 
		LastName: "Agustin", 
		Email: "lailaaa1818@gmail.com", 
		isActive: true,
	}

	fmt.Println(user)
	fmt.Println(user2)
}