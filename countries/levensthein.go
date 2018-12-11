package countries

// import "fmt"

func Lev(a, b []rune) int {
	// Init "multidimensional" array
	table := make([][]int, len(a) + 1)
	for i := range table {
		table[i] = make([]int, len(b) + 1)
	}

	// Init top rows
	for i := range table {
		table[i][0] = i
	}
	for i := range table[0] {
		table[0][i] = i
	}

	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[i]); j++ {
			subCost := 2 // Substitution cost
			if a[i - 1] == b[j - 1] {
				subCost = 0
			}

			// Compare char of string a to char of string b
			// Increase min by 1 if they're not the same
			table[i][j] = findMin(
				table[i - 1][j] + 1, // deletion
				table[i - 1][j - 1] + subCost, // substitution
				table[i][j - 1] + 1) // insertion
		}
	}

	return table[len(table) - 1][len(table[0]) - 1]
}

func findMin(a, b, c int) int {
	min := a

	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	return min
}
