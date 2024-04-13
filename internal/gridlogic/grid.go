package gridlogic

import (
	"sync"
)

// sumSurroundingCells calculates the total score for a specific cell by adding the value of the cell itself and the values of all its directly adjacent cells.
func sumSurroundingCells(grid []int, currentRow, currentCol, gridSize int) int {
	currentCellValue := grid[currentRow*gridSize+currentCol]
	totalSum := currentCellValue

	for deltaRow := -1; deltaRow <= 1; deltaRow++ {
		for deltaCol := -1; deltaCol <= 1; deltaCol++ {
			neighborRow := currentRow + deltaRow
			neighborCol := currentCol + deltaCol
			// Check if the neighbor cell is within the grid boundaries and is not the center cell itself.
			if neighborRow >= 0 && neighborRow < gridSize && neighborCol >= 0 && neighborCol < gridSize && (deltaRow != 0 || deltaCol != 0) {
				neighborValue := grid[neighborRow*gridSize+neighborCol]
				totalSum += neighborValue
			}
		}
	}
	return totalSum
}

// CalculateGridScores computes the scores for all cells in a grid based on their surrounding cells.
func CalculateGridScores(grid []int, gridDimension int) [][]int {
	if len(grid) == 0 || len(grid) != gridDimension*gridDimension {
		return nil // Return nil if the grid is empty or the grid size does not match the expected dimensions.
	}

	scoreMatrix := make([][]int, gridDimension)
	var waitGroup sync.WaitGroup

	for rowIndex := 0; rowIndex < gridDimension; rowIndex++ {
		waitGroup.Add(1)
		go func(rowIndex int) {
			scoreMatrix[rowIndex] = make([]int, gridDimension)
			for colIndex := 0; colIndex < gridDimension; colIndex++ {
				scoreMatrix[rowIndex][colIndex] = sumSurroundingCells(grid, rowIndex, colIndex, gridDimension)
			}
			waitGroup.Done()
		}(rowIndex)
	}
	waitGroup.Wait()
	return scoreMatrix
}
