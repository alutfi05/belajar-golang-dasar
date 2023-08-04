package main

import "fmt"

func main() {
	sentence := printMyResult("Saya sedang")
	fmt.Println(sentence)

	result := add(5, 2)
	fmt.Println(result)
}

func printMyResult(sentence string) string {
	newSentence := sentence + " belajar golang"
	return newSentence
}

func add(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}