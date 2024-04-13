package gridlogic

import (
	"fmt"
	"reflect"
	"testing"
)

// TestSumSurroundingCells tests the sumSurroundingCells function to ensure it calculates correct cell sums.
func TestSumSurroundingCells(t *testing.T) {
	grid := []int{
		4, 2, 3, 2,
		0, 1, 2, 2,
		1, 3, 0, 2,
		2, 0, 1, 5,
	}
	size := 4

	tests := []struct {
		row, col int
		want     int
	}{
		{row: 1, col: 1, want: 16}, // Sum at center cell
		{row: 0, col: 0, want: 7},  // Sum at corner cell
		{row: 3, col: 3, want: 8},  // Sum at another corner cell
	}

	for _, tt := range tests {
		got := sumSurroundingCells(grid, tt.row, tt.col, size)
		if got != tt.want {
			t.Errorf("sumSurroundingCells(%d, %d) = %d; want %d", tt.row, tt.col, got, tt.want)
		}
	}
}

// TestCalculateGridScores tests the grid score calculation for accuracy.
func TestCalculateGridScores(t *testing.T) {
	grid := []int{
		4, 2, 3, 2,
		0, 1, 2, 2,
		1, 3, 0, 2,
		2, 0, 1, 5,
	}
	rowLength := 4
	expected := [][]int{
		{7, 12, 12, 9},
		{11, 16, 17, 11},
		{7, 10, 16, 12},
		{6, 7, 11, 8},
	}

	scores := CalculateGridScores(grid, rowLength)
	if !reflect.DeepEqual(scores, expected) {
		t.Errorf("CalculateGridScores() got %v, want %v", scores, expected)
		fmt.Println("Computed scores:")
		for _, row := range scores {
			fmt.Println(row)
		}
		fmt.Println("Expected scores:")
		for _, row := range expected {
			fmt.Println(row)
		}
	}
}

// TestGetTopScores tests that the highest scores are correctly identified and formatted.
func TestGetTopScores(t *testing.T) {
	scores := [][]int{
		{7, 12, 12, 9},
		{9, 16, 16, 13},
		{9, 17, 13, 11},
		{6, 10, 11, 9},
	}
	topCount := 2
	expected := "(2, 1, 17) (1, 1, 16)"

	result := GetTopScores(scores, topCount)
	if result != expected {
		t.Errorf("GetTopScores() got %v, want %v", result, expected)
	}
}
