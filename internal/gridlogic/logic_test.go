package gridlogic

import (
	"testing"
)

func TestGetTopScores(t *testing.T) {
	scores := [][]int{{7, 12, 12, 9}, {9, 16, 16, 13}, {9, 17, 13, 11}, {6, 10, 11, 9}}
	topCount := 2
	wantResult := "(2, 1, 17)(1, 1, 16)"

	result := GetTopScores(scores, topCount)
	if result != wantResult {
		t.Errorf("GetTopScores() = %v, want %v", result, wantResult)
	}
}
