package board

import "fmt"

// Type Bitboard = uint64
type Bitboard uint64

// Defining square names
const (
	A8 = iota
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)

const (
	White = iota
	Black
)

var SquareToCordinates = [...]string{
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
}

// Print Board
func PrintBitboard(bb Bitboard) {
	fmt.Printf("\n\n")

	// Loop over board ranks
	for rank := 0; rank < 8; rank++ {
		// Loop over board files
		for file := 0; file < 8; file++ {
			// convert file and rank to square index
			square := rank*8 + file

			// Print ranks of the board
			if file == 0 {
				fmt.Printf("%d  ", 8-rank)
			}

			// Print bit state as 1 or 0
			if GetBit(bb, square) != 0 {
				fmt.Printf(" 1 ")
			} else {
				fmt.Printf(" 0 ")
			}
		}

		// Print new line every rank
		fmt.Printf("\n")
	}
	// Print files of the board
	fmt.Printf("\n    a  b  c  d  e  f  g  h\n")
	fmt.Printf("\n    Bitboard: %d\n", bb)
}
