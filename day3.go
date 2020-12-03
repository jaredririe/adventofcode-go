package main

const (
	tree = "#"
	open = "."
)

type Slope struct {
	Right int
	Down  int
}

func TreesEncountered(mapPattern []string, slopes []Slope) int {

	var pattern [][]string
	var treesPerSlope []int

	for _, row := range mapPattern {
		var r []string

		for i := 0; i < len(row); i++ {
			r = append(r, string([]rune(row)[i]))
		}

		pattern = append(pattern, r)
	}

	width := len(pattern[0])
	totalRows := len(pattern)

	for _, slope := range slopes {

		var trees int

		// Start in the top left corner: (0, 0)
		row, col := 0, 0

		for row < totalRows-slope.Down {

			// Move according to the given slope (right X, down Y)
			col += slope.Right
			row += slope.Down

			// Handle repeating pattern
			col %= width

			// See if a tree was encountered
			if pattern[row][col] == tree {
				trees++
			}
		}

		treesPerSlope = append(treesPerSlope, trees)
	}

	// The total is found by multiplying together
	// the number of trees encountered at each slope
	total := 1
	for _, trees := range treesPerSlope {
		total *= trees
	}

	return total
}
