package attacks

import "github.com/theflippantfox/chessengine/packages/board"

// Knight attacks table
var KnightAttacks [64]board.Bitboard

// Generate rook attack table
func MaskKnightAttacks(square int) board.Bitboard {
	// Result attacks bitboard
	attacks := board.Bitboard(0)

	// Piece bitboard
	bitboard := board.Bitboard(0)

	// Set piece on board
	bitboard = board.SetBit(bitboard, square)

	if (bitboard >> 17 & not_h_file) != 0 {
		attacks = (attacks | (bitboard >> board.Bitboard(17)))
	}
	if (bitboard >> 15 & not_a_file) != 0 {
		attacks = (attacks | (bitboard >> board.Bitboard(15)))
	}
	if (bitboard >> 10 & not_gh_file) != 0 {
		attacks = (attacks | (bitboard >> board.Bitboard(10)))
	}
	if (bitboard >> 6 & not_ab_file) != 0 {
		attacks = (attacks | (bitboard >> board.Bitboard(6)))
	}
	if (bitboard << 17 & not_a_file) != 0 {
		attacks = (attacks | (bitboard << board.Bitboard(17)))
	}
	if (bitboard << 15 & not_h_file) != 0 {
		attacks = (attacks | (bitboard << board.Bitboard(15)))
	}
	if (bitboard << 10 & not_ab_file) != 0 {
		attacks = (attacks | (bitboard << board.Bitboard(10)))
	}
	if (bitboard << 6 & not_gh_file) != 0 {
		attacks = (attacks | (bitboard << board.Bitboard(6)))
	}

	return attacks
}
