package main

import (
	"fmt"
	"struct/management"
)

func main() {
	user := management.User{}
	user.ID = 1
	user.FirstName = "Ahmad Lutfi"
	user.LastName = "Rizki Patria"
	user.Email = "aluthfi0522@gmail.com"
	user.IsActive = true

	fmt.Println(user.Display())

	user2 := management.User{
		ID: 2,
		FirstName: "Laila Indah", 
		LastName: "Agustin", 
		Email: "lailaaa1818@gmail.com", 
		IsActive: true,
	}

	users := []management.User{user, user2}

	group := management.Group{
		Name: "Gaming",
		Admin: user,
		Users: users,
		IsAvailable: true,
	}

	group.DisplayGroup()
}