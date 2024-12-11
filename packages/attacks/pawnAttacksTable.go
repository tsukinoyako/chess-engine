package attacks

import "github.com/theflippantfox/chessengine/packages/board"

// Pawn attacks table
var PawnAttacks [2][64]board.Bitboard

// Generate pawn attack table
func MaskPawnAttacks(square, side int) board.Bitboard {
	// Result attacks bitboard
	attacks := board.Bitboard(0)

	// Piece bitboard
	bitboard := board.Bitboard(0)

	// Set piece on board
	bitboard = board.SetBit(bitboard, square)

	// White Pawns
	if side == 0 {
		attacks = not_a_file & (attacks | (bitboard >> board.Bitboard(7)))
		attacks = not_h_file & (attacks | (bitboard >> board.Bitboard(9)))

		// Black pawns
	} else {
		attacks = not_h_file & (attacks | (bitboard << board.Bitboard(7)))
		attacks = not_a_file & (attacks | (bitboard << board.Bitboard(9)))
	}

	return attacks
}
