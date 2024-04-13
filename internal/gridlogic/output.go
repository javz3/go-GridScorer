package gridlogic

import (
	"fmt"
	"sort"
	"strings"
)

// GetTopScores computes and returns a formatted string of the top scores from a grid.
// The top scores are sorted in descending order based on their values.
func GetTopScores(gridScores [][]int, numberOfTopScores int) string {
	var scoreEntries []scoreEntry
	for rowIndex, row := range gridScores {
		for colIndex, score := range row {
			scoreEntries = append(scoreEntries, scoreEntry{rowIndex: rowIndex, colIndex: colIndex, score: score})
		}
	}

	// Sort score entries in descending order by score
	sort.Slice(scoreEntries, func(i, j int) bool {
		return scoreEntries[i].score > scoreEntries[j].score
	})

	var result strings.Builder
	for i := 0; i < numberOfTopScores && i < len(scoreEntries); i++ {
		// Format and append the top score entry to the result string
		formattedScore := fmt.Sprintf("(%d, %d, %d)", scoreEntries[i].rowIndex, scoreEntries[i].colIndex, scoreEntries[i].score)
		if i > 0 {
			result.WriteString(" ") // Add a space between entries for readability
		}
		result.WriteString(formattedScore)
	}

	return result.String()
}
