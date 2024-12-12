package attacks

import "github.com/theflippantfox/chessengine/packages/board"

// Pawn attacks table
var PawnAttacks [2][64]board.Bitboard

// Generate pawn attack table
func MaskPawnAttacks(square, side int) board.Bitboard {
	attacks := board.Bitboard(0)

	// Attack directions for white and black pawns
	directions := [2][]int{
		{-7, -9}, // White: North-East, North-West
		{7, 9},   // Black: South-East, South-West
	}

	// File masks to prevent wrapping
	masks := []board.Bitboard{not_a_file, not_h_file}

	for _, direction := range directions[side] {
		targetSquare := square + direction

		// Ensure the target square is within bounds
		if targetSquare < 0 || targetSquare >= 64 {
			continue
		}

		// Apply file mask to prevent wrapping
		if (direction == -7 || direction == 7) && (masks[1]&(1<<square)) == 0 {
			attacks |= board.Bitboard(1) << targetSquare
		}
		if (direction == -9 || direction == 9) && (masks[0]&(1<<square)) == 0 {
			attacks |= board.Bitboard(1) << targetSquare
		}
	}

	return attacks
}
