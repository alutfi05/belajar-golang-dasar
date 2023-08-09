package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// Method with pointer receiver
func (student *Student) graduate()  {
	student.Name += " A.Md.Kom."
}

func main() {
	student := Student{
		ID:   1,
		Name: "Ahmad Lutfi Rizki Patria",
		GPA:  3.28,
	}

	// Expect : Ahmad Lutfi Rizki Patria
	fmt.Println(student.Name)

	// graduate(&student)
	student.graduate()

	// Expect : Ahmad Lutfi Rizki Patria A.Md.Kom.
	fmt.Println(student.Name)
}

// Function with pointer struct as parameter
// func graduate(student *Student)  {
// 	student.Name += " A.Md.Kom."
// }