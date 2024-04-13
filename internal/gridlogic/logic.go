package gridlogic

import (
	"fmt"
	"sort"
	"strings"
)

// sumSurroundingCells calculates the sum of a cell and its surrounding cells in the grid.
func sumSurroundingCells(grid []int, row, col, size int) int {
	sum := grid[row*size+col] // Start with the value of the cell itself
	for dRow := -1; dRow <= 1; dRow++ {
		for dCol := -1; dCol <= 1; dCol++ {
			// Check if we're out of the bounds of the grid
			neighborRow, neighborCol := row+dRow, col+dCol
			if neighborRow >= 0 && neighborRow < size && neighborCol >= 0 && neighborCol < size {
				// Avoid adding the center cell twice
				if dRow != 0 || dCol != 0 {
					sum += grid[neighborRow*size+neighborCol]
				}
			}
		}
	}
	return sum
}

func CalculateGridScores(grid []int, rowLength int) [][]int {
	if len(grid) == 0 || len(grid) != rowLength*rowLength {
		return nil // handle incorrect grid size or empty grid
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

func GetTopScores(scores [][]int, topCount int) string {
	var scoreEntries []struct {
		row, col, score int
	}

	for row := range scores {
		for col, score := range scores[row] {
			scoreEntries = append(scoreEntries, struct {
				row, col, score int
			}{row, col, score})
		}
	}

	sort.Slice(scoreEntries, func(i, j int) bool {
		if scoreEntries[i].score == scoreEntries[j].score {
			if scoreEntries[i].row == scoreEntries[j].row {
				return scoreEntries[i].col < scoreEntries[j].col
			}
			return scoreEntries[i].row < scoreEntries[j].row
		}
		return scoreEntries[i].score > scoreEntries[j].score
	})

	var result strings.Builder
	for i := 0; i < topCount && i < len(scoreEntries); i++ {
		entry := scoreEntries[i]
		result.WriteString(fmt.Sprintf("(%d, %d, %d)", entry.row, entry.col, entry.score))
	}

	return result.String()
}
