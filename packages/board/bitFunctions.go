package board

// Bit Functions
func SetBit(bb Bitboard, square int) Bitboard { return (bb | (Bitboard(1) << square)) }
func GetBit(bb Bitboard, square int) Bitboard { return (bb & (Bitboard(1) << square)) }
func ClearBit(bb Bitboard, square int) Bitboard {
	if GetBit(bb, square) != 0 {
		return (bb ^ (Bitboard(1) << square))
	}
	return bb
}
