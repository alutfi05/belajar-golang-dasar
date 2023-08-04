package main

import "fmt"

func main() {
	sentence := printMyResult("Saya sedang")
	fmt.Println(sentence)

	result := add(5, 2)
	fmt.Println(result)

	area, around := rectangle(5, 2);
	fmt.Println("Luas dan keliling persegi panjang")
	fmt.Println("Luas :", area)
	fmt.Println("Keliling :", around)
	
	// luas, _ := rectangle(5, 2);
	// fmt.Println("Luas :", luas)

	luas, keliling := square(3, 3)
	fmt.Println("Luas dan keliling persegi")
	fmt.Println("Luas :", luas)
	fmt.Println("Keliling :", keliling)
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
func rectangle(length, width int) (int, int) {
	area := length * width
	around := 2 * (length + width)

	return area, around
}

// Predefined return value
func square(sideOne, sideTwo int) (area, around int) {
	area = sideOne * sideTwo
	around = 4 * sideOne

	return
}