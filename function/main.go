package main

import "fmt"

func main() {
	sentence := printMyResult("Saya sedang")
	fmt.Println(sentence)

	result := add(5, 2)
	fmt.Println(result)

	luas, keliling := rectangle(5, 2);
	fmt.Println("Luas :", luas)
	fmt.Println("Keliling :", keliling)
	
	// luas, _ := rectangle(5, 2);
	// fmt.Println("Luas :", luas)
}

// Basic function
func printMyResult(sentence string) string {
	newSentence := sentence + " belajar golang"
	return newSentence
}

// Advanced function
func add(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}

// Function with multiple return
func rectangle(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}