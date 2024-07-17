// Package chessboard provides functionality to represent and work with a chessboard.
package chessboard

// File represents a column on a chessboard, with each boolean value indicating
// whether a square in that column is occupied by a piece or not.
type File []bool

// Chessboard represents a standard 8x8 chessboard using a map where keys are
// file names ("A" to "H") and values are File types representing the columns.
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, lb string) (c int) {
	for _, v := range cb[lb] {
		if v {
			c++
		}
	}
	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (c int) {
	if rank < 1 || rank > 8 {
		return
	}
	for _, f := range cb {
		if f[rank-1] {
			c++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (c int) {
	for _, f := range cb {
		for range f {
			c++
		}
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (c int) {
	for _, f := range cb {
		for _, v := range f {
			if v {
				c++
			}
		}
	}
	return
}
