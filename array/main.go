package main

import "fmt"

func main() {
	// var languages [5]string
	// languages[0] = "Node.js"
	// languages[1] = "Typescript"
	// languages[2] = "Javascript"
	// languages[3] = "Golang"
	// languages[4] = "Python"

	// languages := [5]string {
	// 	"Node.js", 
	// 	"Typescript", 
	// 	"Javascript", 
	// 	"Golang", 
	// 	"Python",
	// }

	languages := [...]string {
		"Node.js", 
		"Typescript", 
		"Javascript", 
		"Golang", 
		"Python",
		"Java",
	}

	fmt.Println(languages)

	length := len(languages)
	fmt.Println(length)

	for index, languange := range languages {
		fmt.Println("Index :", index, "language :", languange)
	}
}