package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func main() {
	student := Student{
		ID:   1,
		Name: "Ahmad Lutfi Rizki Patria",
		GPA:  3.28,
	}

	// Expect : Ahmad Lutfi Rizki Patria
	fmt.Println(student.Name)

	graduate(&student)

	// Expect : Ahmad Lutfi Rizki Patria A.Md.Kom.
	fmt.Println(student.Name)
}

func graduate(student *Student)  {
	student.Name += " A.Md.Kom."
}