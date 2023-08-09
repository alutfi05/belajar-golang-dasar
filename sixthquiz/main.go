package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

// Add game
func (gamer *Gamer) AddGame(name string)  {
	gamer.Games = append(gamer.Games, name)
}

func main() {
	fmt.Println("Quiz 6 : create method AddGame with pointer receiver and append to slice games")

	gamer := Gamer{
		Name:  "Ahmad Lutfi Rizki Patria",
	}

	// Expect :
	// Dota 2
	// League of Legends
	// Mobile Legends

	gamer.AddGame("Dota 2")
	gamer.AddGame("League of Legends")
	gamer.AddGame("Mobile Legends")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}