package main

import (
	"github.com/tsukinoyako/chessengine/pkg/attacks"
	"github.com/tsukinoyako/chessengine/pkg/board"
)

/*
   "a8" "b8" "c8" "d8" "e8" "f8" "g8" "h8"
   "a7" "b7" "c7" "d7" "e7" "f7" "g7" "h7"
   "a6" "b6" "c6" "d6" "e6" "f6" "g6" "h6"
   "a5" "b5" "c5" "d5" "e5" "f5" "g5" "h5"
   "a4" "b4" "c4" "d4" "e4" "f4" "g4" "h4"
   "a3" "b3" "c3" "d3" "e3" "f3" "g3" "h3"
   "a2" "b2" "c2" "d2" "e2" "f2" "g2" "h2"
   "a1" "b1" "c1" "d1" "e1" "f1" "g1" "h1"
*/

func main() {
	// Init Leaper Attacks
	attacks.InitLeaperAttacks()

	// Init occupancy board
	block := board.Bitboard(0)
	block = board.SetBit(block, board.D7)
	block = board.SetBit(block, board.C4)
	block = board.SetBit(block, board.D3)
	block = board.SetBit(block, board.G4)

	//board.FENParser("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
}