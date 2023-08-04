package main

import "fmt"

func main() {
	// Quiz 3
	// Hitung rata rata dari array scores
	// Filter scores >= 90 dan append ke slice goodScores

	// Hitung rata rata
	var total int
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	for _, score := range scores {
		total += score
	}

	length := len(scores)
	average := float64(total) / float64(length)

	fmt.Println("Average :", average)


	// Filter scores >= 90, append kedalam goodScores
	var goodScores []int
	scores = [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println("Good Scores :", goodScores)
}