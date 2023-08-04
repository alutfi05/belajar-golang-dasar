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
		"javascript": "is very usefull",
	}

	fmt.Println(myMap)
	fmt.Println(myMap["typescript"])
	fmt.Println(myMap["golang"])

	for key, value := range myMap {
		fmt.Println("Key :", key, ", value :", value)
	}

	fmt.Println("==============")

	// Delete item in map
	delete(myMap, "javascript")

	for key, value := range myMap {
		fmt.Println("Key :", key, ", value :", value)
	}

	// Checking key is exist
	value, isAvailable := myMap["PHP"]
	fmt.Println(value)
	fmt.Println(isAvailable)
}