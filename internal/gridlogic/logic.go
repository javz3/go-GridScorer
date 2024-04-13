package gridlogic

import (
	"fmt"
	"sort"
)

// sumSurroundingCells calculates the sum of a cell and its surrounding cells in the grid.
func sumSurroundingCells(grid []int, row, col, size int) int {
	sum := grid[row*size+col] // Start with the value of the cell itself

	for dRow := -1; dRow <= 1; dRow++ {
		for dCol := -1; dCol <= 1; dCol++ {
			neighborRow, neighborCol := row+dRow, col+dCol
			if neighborRow >= 0 && neighborRow < size && neighborCol >= 0 && neighborCol < size {
				contribution := grid[neighborRow*size+neighborCol]
				if dRow != 0 || dCol != 0 { // Avoid adding the center cell twice
					sum += contribution
				}
			}
		}
	}
	return sum
}

// CalculateGridScores processes the given grid and computes the sum of each cell with its neighbors.
func CalculateGridScores(grid []int, rowLength int) [][]int {
	if len(grid) == 0 || len(grid) != rowLength*rowLength {
		return nil // Handle incorrect grid size or empty grid
	}

	scores := make([][]int, rowLength)
	for row := 0; row < rowLength; row++ {
		scores[row] = make([]int, rowLength)
		for col := 0; col < rowLength; col++ {
			scores[row][col] = sumSurroundingCells(grid, row, col, rowLength)
		}
	}
	return scores
}

func GetTopScores(scores [][]int, countOfHighScores int) string {
	var entries []scoreEntry
	for i, row := range scores {
		for j, score := range row {
			entries = append(entries, scoreEntry{x: i, y: j, score: score})
		}
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].score > entries[j].score // Sort descending by score
	})

	result := ""
	for i := 0; i < countOfHighScores; i++ {
		if i < len(entries) {
			result += fmt.Sprintf("(%d, %d, %d)", entries[i].x, entries[i].y, entries[i].score)
		}
	}
	return result
}
