package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	isActive  bool
}

type Group struct {
	Name 		string
	Admin 		User
	Users 		[]User
	isAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName, user.LastName)
	}
}

func main() {
	fmt.Println("Quiz 5 : change function display group into a method")

	user1 := User{
		ID: 1,
		FirstName: "Ahmad Lutfi",
		LastName: "Rizki Patria",
		Email: "aluthfi0522@gmail.com",
		isActive: true,
	}

	user2 := User{
		ID: 2,
		FirstName: "Laila Indah", 
		LastName: "Agustin", 
		Email: "lailaaa1818@gmail.com", 
		isActive: true,
	}

	users := []User{user1, user2}
	
	group := Group{
		Name: "Create display group method",
		Admin: user1,
		Users: users,
		isAvailable: true,
	}

	group.DisplayGroup()
}