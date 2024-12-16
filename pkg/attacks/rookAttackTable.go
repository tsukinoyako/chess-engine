package attacks

import (
	"github.com/tsukinoyako/chessengine/pkg/board"
)

// Generate rook attacks
func MaskRookAttacks(square int) board.Bitboard {
	attacks := board.Bitboard(0)

	// Target rank and file
	tr, tf := square/8, square%8

	// Helper to calculate attacks in a single direction
	traverse := func(dr, df int) {
		r, f := tr+dr, tf+df
		for r >= 1 && r <= 6 && f >= 1 && f <= 6 {
			// Set the attack bit for the current position
			attacks |= (board.Bitboard(1) << (r*8 + f))
			r += dr
			f += df
		}
	}

	// Traverse all four orthogonal directions
	traverse(1, 0)  // North
	traverse(-1, 0) // South
	traverse(0, 1)  // East
	traverse(0, -1) // West

	return attacks
}

// Generate rook attacks on the fly
func GenerateRookAttacksOnTheFly(square int, block board.Bitboard) board.Bitboard {
	attacks := board.Bitboard(0)

	// Target rank and file
	tr, tf := square/8, square%8

	// Helper to calculate attacks in a single direction
	traverse := func(dr, df int) {
		r, f := tr+dr, tf+df
		for r >= 0 && r < 8 && f >= 0 && f < 8 {

			// Set the attack bit for the current position
			attacks |= (board.Bitboard(1) << (r*8 + f))

			// If the current position is blocked, stop the attack
			if (board.Bitboard(1) << (r*8 + f) & block) != 0 {
				break
			}

			r += dr
			f += df
		}
	}

	// Traverse all four orthogonal directions
	traverse(1, 0)  // North
	traverse(-1, 0) // South
	traverse(0, 1)  // East
	traverse(0, -1) // West

	return attacks
}
