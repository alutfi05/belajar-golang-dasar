package main

import "fmt"

func main() {
	var grade string
	score := 80

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else {
		grade = "D"
	}

	fmt.Println(grade)
}