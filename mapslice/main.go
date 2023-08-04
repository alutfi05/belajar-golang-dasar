package main

import "fmt"

func main() {
	students := []map[string]string{
		{
			"name":  "Ahmad Lutfi",
			"score": "A",
		},
		{
			"name":  "Rizki Patria",
			"score": "B",
		},
		{
			"name":  "Ahmad Lutfi Rizki Patria",
			"score": "C",
		},
	}

	fmt.Println(students)

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}
}