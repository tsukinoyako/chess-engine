package main

import (
	"github.com/tsukinoyako/chessengine/pkg/attacks"
	"github.com/tsukinoyako/chessengine/pkg/board"
)

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
