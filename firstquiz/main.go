package main

import (
	"firstquiz/calculation"
	"fmt"
)

func main() {
	fmt.Println("Quiz 1 : Multiply two number");

	result := calculation.Multiply(5,3)
	fmt.Println(result)
}