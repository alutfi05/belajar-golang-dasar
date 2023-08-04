package main

import "fmt"

func main() {
	// gamingConsoles := []string{
	// 	"PlayStation4",
	// 	"Nintendo Switch",
	// 	"Xbox One",
	// }

	var gamingConsoles []string
	gamingConsoles = append(gamingConsoles, "PlayStation4")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	gamingConsoles = append(gamingConsoles, "Xbox One")

	fmt.Println(gamingConsoles)

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}