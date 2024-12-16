package attacks

import (
	"github.com/tsukinoyako/chessengine/pkg/board"
)

// Generate Bishop Attacks
func MaskBishopAttacks(square int) board.Bitboard {
	attacks := board.Bitboard(0)

	// Target rank and file
	tr, tf := square/8, square%8

	// Helper to calculate diagonal attacks
	traverse := func(dr, df int) {
		r, f := tr+dr, tf+df
		for r >= 1 && r <= 6 && f >= 1 && f <= 6 {
			attacks |= (board.Bitboard(1) << (r*8 + f))

			r += dr
			f += df
		}
	}

	// Traverse all four diagonals
	traverse(1, 1)   // North-East
	traverse(-1, 1)  // South-East
	traverse(1, -1)  // North-West
	traverse(-1, -1) // South-West

	return attacks
}

// Generate bishop attack moves on the fly, considering blockers
func GenerateBishopAttacksOnTheFly(square int, block board.Bitboard) board.Bitboard {
	// Initialize result bitboard
	attacks := board.Bitboard(0)

	// Target rank and file
	tr, tf := square/8, square%8

	// Helper to generate diagonal attacks in specified direction
	traverse := func(dr, df int) {
		r, f := tr+dr, tf+df
		for r >= 0 && r < 8 && f >= 0 && f < 8 {
			// Set the attack bit for the current position
			attacks |= (board.Bitboard(1) << (r*8 + f))

			// If the current position is blocked, stop the attack
			if (board.Bitboard(1) << (r*8 + f) & block) != 0 {
				break
			}

			// Move to the next position in the same direction
			r += dr
			f += df
		}
	}

	// Traverse all four diagonals
	traverse(1, 1)   // North-East
	traverse(-1, 1)  // South-East
	traverse(1, -1)  // North-West
	traverse(-1, -1) // South-West

	return attacks
}
