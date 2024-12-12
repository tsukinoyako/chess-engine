package attacks

import "github.com/theflippantfox/chessengine/packages/board"

// Knight attacks table
var KnightAttacks [64]board.Bitboard

// Generate knight attack table
func MaskKnightAttacks(square int) board.Bitboard {
	attacks := board.Bitboard(0)

	// Knight move offsets
	moves := []int{
		17, 15, 10, 6, // Moves in positive directions
		-17, -15, -10, -6, // Moves in negative directions
	}

	// File restrictions for each knight move
	masks := []board.Bitboard{
		not_h_file, not_a_file, not_gh_file, not_ab_file, // Positive moves
		not_a_file, not_h_file, not_ab_file, not_gh_file, // Negative moves
	}

	for i, move := range moves {
		targetSquare := square + move

		// Ensure the move stays within bounds
		if targetSquare < 0 || targetSquare >= 64 {
			continue
		}

		// Apply file restriction to prevent wrapping
		if (board.Bitboard(1)<<square)&masks[i] != 0 {
			attacks |= board.Bitboard(1) << targetSquare
		}
	}

	return attacks
}
