package main

import (
	"errors"
	"fmt"
)

func main() {
	// Quiz 4
	// Create function sum
	// Result : 39
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)

	fmt.Println("Result sum :", total)

	// Create function calculate
	// Result : 12
	resultPlus, err := calculate(10, 2, "+");
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result (+) :", resultPlus)

	// Result : 8
	resultMin, err := calculate(10, 2, "-");
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result (-) :", resultMin)

	// Result : 20
	resultMultiply, err := calculate(10, 2, "*");
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result (*) :", resultMultiply)

	// Result : 5
	resultDistribution, err := calculate(10, 2, "/");
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result (/) :", resultDistribution)

	// Error Message : invalid operation, please input (+, -, * or /)
	// Result : 0
	result, err := calculate(10, 2, "=");
	if err != nil {
		fmt.Println("Error message :", err.Error())
	}

	fmt.Println("Result :", result)
}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}

	return total
}

func calculate(numberOne, numberTwo int, operation string) (int, error) {
	var result int
	var errorMessage error

	if operation == "+" {
		result = numberOne + numberTwo
	} else if operation == "-" {
		result = numberOne - numberTwo
	} else if operation == "*" {
		result = numberOne * numberTwo
	} else if operation == "/" {
		result = numberOne / numberTwo
	} else {
		errorMessage = errors.New("invalid operation, please input (+, -, * or /)")
	}

	return result, errorMessage
}