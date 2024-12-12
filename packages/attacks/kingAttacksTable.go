package attacks

import "github.com/theflippantfox/chessengine/packages/board"

// King attacks table
var KingAttacks [64]board.Bitboard

// Generate king attack table
func MaskKingAttacks(square int) board.Bitboard {
	attacks := board.Bitboard(0)

	// Directions for king movement
	directions := []int{
		8,  // North
		9,  // North-East
		7,  // North-West
		1,  // East
		-8, // South
		-9, // South-West
		-7, // South-East
		-1, // West
	}

	// Iterate through all directions
	for _, direction := range directions {
		targetSquare := square + direction
		if targetSquare < 0 || targetSquare >= 64 {
			continue // Skip out-of-bounds squares
		}

		shift := board.Bitboard(1) << targetSquare

		// Handle file wrapping
		if direction == -9 || direction == 7 || direction == -1 {
			shift &= not_h_file
		}
		if direction == -7 || direction == 9 || direction == 1 {
			shift &= not_a_file
		}

		// Add the attack bit
		attacks |= shift
	}

	return attacks
}
