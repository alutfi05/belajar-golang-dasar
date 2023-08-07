package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	isActive  bool
}

// Method User Struct
func (user User) display() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

// Embedded Struct
type Group struct {
	Name 		string
	Admin 		User
	Users 		[]User
	isAvailable bool
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

	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user2)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

	// Embedded Struct
	users := []User{user, user2}

	group := Group{
		Name: "Gaming",
		Admin: user,
		Users: users,
		isAvailable: true,
	}

	displayGroup(group)

	// Method Struct
	resultDisplay1 := user.display()
	resultDisplay2 := user2.display()

	fmt.Println(resultDisplay1)
	fmt.Println(resultDisplay2)
}

// Struct as parameter 
func displayUser(user User) string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

func displayGroup(group Group) {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}