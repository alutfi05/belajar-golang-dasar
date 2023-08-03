package main

import (
	"fmt"
)

func main() {
	// Quiz 2
	// 1. Mencetak huruf yang index nya genap
	// 2. Mencetak huruf vocal : a, i, u, e, o

	fmt.Println("Quiz 2 : Mencetak huruf yang index nya genap dan mencetak huruf vocal")

	title := "Golang the best language"

	for index, letter := range title {
		letterString := string(letter)
		
		if index % 2 == 0 {
			fmt.Println("Index :", index, "huruf :", letterString)
		}
	}

	for _, letter := range title {
		letterString := string(letter)

		// switch letterString {
		// case "a":
		// 	fmt.Println("vocal :", string(letter))
		// case "i":
		// 	fmt.Println("vocal :", string(letter))
		// case "u":
		// 	fmt.Println("vocal :", string(letter))
		// case "e":
		// 	fmt.Println("vocal :", string(letter))
		// case "o":
		// 	fmt.Println("vocal :", string(letter))
		// }

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("vocal :", letterString)
			
		}
	}

}