package main

import "fmt"

func main() {
	numberA := 5

	// Pointer / Reference -> &
	// Untuk mengakses alamat memory yang disimpan
	numberB := &numberA

	fmt.Println("numberA :", numberA)
	fmt.Println("numberB, pointer -> &numberA :", numberB)

	// Deferencing -> *
	// Untuk mengakses nilai dari memory yang di simpan di Pointer 
	fmt.Println("numberB, deferencing -> *numberB :", *numberB)

	// Ketika numberB diubah numberA akan ikut diubah, karena mengacu ke alamat memory yang sama
	*numberB = 10
	fmt.Println("numberB :", *numberB)
	fmt.Println("numberA :", numberA)

	// Advanced pointer
	var numberOne int = 10
	var numberTwo *int = &numberOne

	fmt.Println("numberOne :", numberOne)
	fmt.Println("numberTwo reference (&):", numberTwo)
	fmt.Println("numberTwo dereference (*):", *numberTwo)

	numberOne = 50
	fmt.Println("Seletah numberOne diubah menjadi 50")
	fmt.Println("numberOne :", numberOne)
	fmt.Println("numberTwo reference (&) :", numberTwo)
	fmt.Println("numberTwo dereference (*):", *numberTwo)

	// Studi kasus
	fmt.Println("Studi Kasus")
	number := 5
	// fmt.Println("Alamat memory :", &number)
	fmt.Println("Old number :", number)
	
	change(&number, 100)
	// fmt.Println("Alamat memory :", &number)
	fmt.Println("New number :", number)
}

func change(old *int, new int) {
	*old = new
	// fmt.Println("Alamat memory :", old)
	// fmt.Println("function :", *old)
}