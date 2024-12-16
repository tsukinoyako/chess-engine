package attacks

import "github.com/tsukinoyako/chessengine/pkg/board"

// To avoid jumping across the board
const not_a_file board.Bitboard = 18374403900871474942
const not_h_file board.Bitboard = 9187201950435737471
const not_ab_file board.Bitboard = 18229723555195321596
const not_gh_file board.Bitboard = 4557430888798830399

// Init leaper pieces attack tables
func InitLeaperAttacks() {
	for square := 0; square < 64; square++ {
		PawnAttacks[board.White][square] = MaskPawnAttacks(square, board.White)
		PawnAttacks[board.Black][square] = MaskPawnAttacks(square, board.Black)

		KnightAttacks[square] = MaskKnightAttacks(square)
		KingAttacks[square] = MaskKingAttacks(square)
	}
}
