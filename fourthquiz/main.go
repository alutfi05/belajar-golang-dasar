package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)

	fmt.Println("Result sum :", total)

	resultPlus, err := calculate(10, 2, "+");
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result (+) :", resultPlus)

	resultMin, err := calculate(10, 2, "-");
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result (-) :", resultMin)

	resultMultiply, err := calculate(10, 2, "*");
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result (*) :", resultMultiply)

	resultDistribution, err := calculate(10, 2, "/");
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Result (/) :", resultDistribution)

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