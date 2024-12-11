package board

import "github.com/theflippantfox/chessengine/packages/pieces"

func FENParser(fen string) {
	bitboard := Bitboard(0)
	position := 0
	for _, char := range fen {

		if char == '/' {
			continue
		}

		if char >= '1' && char <= '8' {
			for i := 1; i <= int(char-'0'); i++ {
				position += 1
			}
		} else {
			bitboard = SetBit(bitboard, position)
			position += 1
		}
	}
	PrintBitboard(bitboard)
}

func charToPiece(char string) int {
	runeChar := []rune(char)[0]
	return pieces.Pieces[runeChar]
}
