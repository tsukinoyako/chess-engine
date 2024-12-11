package attacks

import "github.com/theflippantfox/chessengine/packages/board"

// Knight attacks table
var KingAttacks [64]board.Bitboard

// Generate rook attack table
func MaskKingAttacks(square int) board.Bitboard {
	// Result attacks bitboard
	attacks := board.Bitboard(0)

	// Pie0ce bitboard
	bitboard := board.Bitboard(0)

	// Set piece on board
	bitboard = board.SetBit(bitboard, square)

	if (bitboard >> 8) != 0 {
		attacks = (attacks | (bitboard >> board.Bitboard(8)))
	}
	if (bitboard >> 9 & not_h_file) != 0 {
		attacks = (attacks | (bitboard >> board.Bitboard(9)))
	}
	if (bitboard >> 7 & not_a_file) != 0 {
		attacks = (attacks | (bitboard >> board.Bitboard(7)))
	}
	if (bitboard >> 1 & not_h_file) != 0 {
		attacks = (attacks | (bitboard >> board.Bitboard(1)))
	}
	if (bitboard << 8) != 0 {
		attacks = (attacks | (bitboard << board.Bitboard(8)))
	}
	if (bitboard << 9 & not_a_file) != 0 {
		attacks = (attacks | (bitboard << board.Bitboard(9)))
	}
	if (bitboard << 7 & not_h_file) != 0 {
		attacks = (attacks | (bitboard << board.Bitboard(7)))
	}
	if (bitboard << 1 & not_a_file) != 0 {
		attacks = (attacks | (bitboard << board.Bitboard(1)))
	}

	return attacks
}
