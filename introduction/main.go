package main

// Noted
// Cara import : import "namaModul/namaPackage", contoh -> import "introduction/calculation"
// Cara pemakaian : calculation.Add(1, 2)

import (
	"fmt"
	"introduction/calculation"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Halo, belajar golang")

	result := calculation.Add(1,2)
	fmt.Println(result)
}