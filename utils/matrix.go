package utils

// IsAdjacence determine if is a square matrix or not
func IsAdjacence(matrix [][]uint8) bool {
	for line := range matrix {
		if len(matrix[line]) != len(matrix) {
			return false
		}
	}
	return true
}
