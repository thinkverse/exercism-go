package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	squares := 0

	for _, square := range cb[rank] {
		if square {
			squares++
		}
	}

	return squares
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	squares := 0

	for _, rank := range cb {
		for index, square := range rank {
			if index == file-1 && square {
				squares++
			}
		}
	}

	return squares
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	squares := 0

	for _, rank := range cb {
		for range rank {
			squares++
		}
	}

	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	squares := 0

	for rank := range cb {
		squares += CountInRank(cb, rank)
	}

	return squares
}
