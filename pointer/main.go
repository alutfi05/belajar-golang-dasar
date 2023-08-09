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
}