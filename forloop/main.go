package main

import "fmt"

func main() {
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya sedang belajar Golang :", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Saya sedang belajar Golang :", i)
	// 	i++
	// }

	title := "Belajar golang dasar"
	for index, letter := range title {
		fmt.Println("Index :", index, " letter :", string(letter))
	}

	// Jika tidak membutuhkan index : bisa menggunakan _
	for _, letter := range title {
		fmt.Println("letter :", string(letter))
	}
}