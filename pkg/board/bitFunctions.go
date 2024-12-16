package board

// Set a bit as ON in the bitboard
func SetBit(bb Bitboard, square int) Bitboard { return (bb | (Bitboard(1) << square)) }

// Gets the bit from the bitboard
func GetBit(bb Bitboard, square int) Bitboard { return (bb & (Bitboard(1) << square)) }

// Set the bit to OFF in the bitboard
func ClearBit(bb Bitboard, square int) Bitboard {
	if GetBit(bb, square) != 0 {
		return (bb ^ (Bitboard(1) << square))
	}
	return bb
}

// Counts the ON bits in the bitboard
func CountBits(bb Bitboard) int {
	count := 0

	for bb > 0 {
		count++

		bb = bb & (bb - 1)
	}

	return count
}
