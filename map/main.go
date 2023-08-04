package main

import "fmt"

func main() {
	// var myMap map[string]int

	// myMap = map[string]int{}

	// myMap["typescript"] = 9
	// myMap["javascript"] = 9
	// myMap["golang"] = 8

	myMap := map[string]string{
		"typescript": "is javascript superset",
		"golang": "is super fast",
	}

	fmt.Println(myMap)
	fmt.Println(myMap["typescript"])
	fmt.Println(myMap["golang"])
}