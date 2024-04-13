package main

import (
	"fmt"
	"gridscorer/internal/gridlogic"
)

func main() {
	grid := []int{4, 2, 3, 2, 0, 1, 2, 2, 1, 3, 0, 2, 2, 0, 1, 5}
	rowLength := 4
	topScores := 2

	scores := gridlogic.CalculateGridScores(grid, rowLength)
	fmt.Println(gridlogic.GetTopScores(scores, topScores))
}
