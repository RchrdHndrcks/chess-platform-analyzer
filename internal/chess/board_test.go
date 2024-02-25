package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeMove(t *testing.T) {
	board := NewBoard()

	board.MakeMove("e2e4")

	assert.Equal(t, Piece(""), board.Board[6][4])
	assert.Equal(t, Piece("P"), board.Board[4][4])
	assert.Equal(t, 1, board.MovesCount)
	assert.Equal(t, "e3", board.InPasantSquare)
	assert.Equal(t, "b", board.Turn)
	assert.Equal(t, []string{"e2e4"}, board.MovesHistory)
}

func Test_MakeMove_Castle(t *testing.T) {
	board := Board{}
	board.TranslateFEN("4r3/8/8/8/8/8/8/4K2R w K - 0 1")

	board.MakeMove("e1g1")

	assert.Equal(t, []string{"e1g1"}, board.MovesHistory)
	assert.Equal(t, []Piece{"", "", "", "", "", "R", "K", ""}, board.Board[7])
}

func Test_AvailableLegalMoves_InitialPosition(t *testing.T) {
	board := NewBoard()

	board.MakeMove("e2e4")
	legalMovesForBlack := board.AvailableLegalMoves()

	assert.ElementsMatch(t, []string{"a7a6", "a7a5", "b7b6", "b7b5", "c7c6", "c7c5", "d7d6", "d7d5",
		"e7e6", "e7e5", "f7f6", "f7f5", "g7g6", "g7g5", "h7h6", "h7h5", "b8a6", "b8c6", "g8h6", "g8f6"},
		legalMovesForBlack)
	assert.Equal(t, "b", board.Turn)
}

func Test_AvailableLegalMoves_CheckedKing(t *testing.T) {
	board := Board{}
	board.TranslateFEN("4k3/4Q3/8/8/8/8/8/8 b - - 0 1")

	legalMoves := board.AvailableLegalMoves()

	assert.Len(t, legalMoves, 1)
	assert.Equal(t, "e8e7", legalMoves[0])
}

func Test_AvailableLegalMoves_CheckedKingWithSameColorRookInBoard(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/1b6/K7/8/1R6/8/8/8 w - - 0 1")

	legalMoves := board.AvailableLegalMoves()

	assert.ElementsMatch(t, []string{"a6b7", "a6a7", "a6b6", "a6b5", "a6a5", "b4b7"}, legalMoves)
}

func Test_AvailableLegalMoves_CheckedKingWithCastleOpion(t *testing.T) {
	board := Board{}
	board.TranslateFEN("1k6/8/8/8/8/6b1/8/R3K2R w - KQ 0 1")

	legalMoves := board.AvailableLegalMoves()

	assert.ElementsMatch(t, []string{"e1d1", "e1d2", "e1e2", "e1f1"}, legalMoves)
}

func Test_AvailableLegalMoves_BishopInKingCastleWay(t *testing.T) {
	board := Board{}
	board.TranslateFEN("1r6/8/8/8/8/7b/n7/R3K2R w - KQ 0 1")

	legalMoves := board.AvailableLegalMoves()

	assert.ElementsMatch(t, []string{"a1a2", "a1b1", "a1c1", "a1d1", "e1d1", "e1d2", "e1e2", "e1f2",
		"h1f1", "h1g1", "h1h2", "h1h3"}, legalMoves)
}

func Test_AvailableLegalMoves_Checkmate(t *testing.T) {
	board := Board{}
	board.TranslateFEN("k7/1Q6/2K5/8/8/8/8/8 b - - 0 1")

	legalMoves := board.AvailableLegalMoves()

	assert.Nil(t, legalMoves)
}

func Test_AvailableLegalMoves_Stalemate(t *testing.T) {
	board := Board{}
	board.TranslateFEN("k7/1R6/2K5/8/8/8/8/8 b - - 0 1")

	legalMoves := board.AvailableLegalMoves()

	assert.Len(t, legalMoves, 0)
}

func Test_availableMoves_OnlyPiecesWithDifferentColorOfTurn(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/8/8/4b3/8/8/8/8 w - - 0 1")

	moves := board.availableMoves()

	assert.Nil(t, moves)
}

func Test_availableMoves_Bishop(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/8/8/4B3/8/8/8/8 w - - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"e5f6", "e5g7", "e5h8", "e5f4", "e5g3", "e5h2",
		"e5d4", "e5c3", "e5b2", "e5a1", "e5d6", "e5c7", "e5b8"}, moves)
}

func Test_availableMoves_BishopWithOpponentPiece(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/6r1/8/4B3/8/8/8/8 w - - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"e5f6", "e5g7", "e5f4", "e5g3", "e5h2",
		"e5d4", "e5c3", "e5b2", "e5a1", "e5d6", "e5c7", "e5b8"}, moves)
}

func Test_availableMoves_BishopWithSameColorPiece(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/6R1/8/4B3/8/8/8/8 w - - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"e5f6", "e5f4", "e5g3", "e5h2",
		"e5d4", "e5c3", "e5b2", "e5a1", "e5d6", "e5c7", "e5b8", "g7g1",
		"g7g2", "g7g3", "g7g4", "g7g5", "g7g6", "g7g8", "g7a7", "g7b7",
		"g7c7", "g7d7", "g7e7", "g7f7", "g7h7"}, moves)
}

func Test_availableMoves_Rook(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/8/8/8/8/8/6r1/8 b - - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"g2a2", "g2b2", "g2c2", "g2d2", "g2e2", "g2f2",
		"g2h2", "g2g1", "g2g3", "g2g4", "g2g5", "g2g6", "g2g7", "g2g8"}, moves)
}

func Test_availableMoves_Queen(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/8/8/4Q3/8/8/8/8 w - - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"e5f6", "e5g7", "e5h8", "e5f4", "e5g3", "e5h2",
		"e5d4", "e5c3", "e5b2", "e5a1", "e5d6", "e5c7", "e5b8", "e5e1", "e5e2", "e5e3",
		"e5e4", "e5e6", "e5e7", "e5e8", "e5a5", "e5b5", "e5c5", "e5d5", "e5f5", "e5g5", "e5h5"}, moves)
}

func Test_availableMoves_King(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/8/8/8/8/8/2K5/8 w - - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"c2b1", "c2b2", "c2b3", "c2c1", "c2c3", "c2d1", "c2d2", "c2d3"}, moves)
}

func Test_availableMoves_KingCastles(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/8/8/8/8/8/8/R3K2R w KQ - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"e1f1", "e1f2", "e1e2", "e1d2", "e1d1", "e1g1", "e1c1", "a1b1",
		"a1c1", "a1d1", "a1a2", "a1a3", "a1a4", "a1a5", "a1a6", "a1a7", "a1a8", "h1g1", "h1f1", "h1h2",
		"h1h3", "h1h4", "h1h5", "h1h6", "h1h7", "h1h8"}, moves)
}

func Test_availableMoves_Knight(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/8/8/8/3N5/8/8/8 w - - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"d4c6", "d4e6", "d4f5", "d4f3", "d4e2", "d4c2", "d4b5", "d4b3"}, moves)
}

func Test_availableMoves_Pawn(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/7P/8/4p3/4P3/P5p1/1P3P2/8 w - - 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"a3a4", "b2b3", "b2b4", "f2f3", "f2f4", "f2g3", "h7h8Q", "h7h8R",
		"h7h8N", "h7h8B"}, moves)
}

func Test_availableMoves_PawnInPassant(t *testing.T) {
	board := Board{}
	board.TranslateFEN("8/8/8/4pP2/8/8/P7/8 w - e6 0 1")

	moves := board.availableMoves()

	assert.ElementsMatch(t, []string{"a2a3", "a2a4", "f5f6", "f5e6"}, moves)
}

func Test_FEN(t *testing.T) {
	board := Board{
		Board: [][]Piece{
			{BRook, "", "", "", BRook, "", BKing, ""},
			{BPawn, BPawn, "", "", "", BKnight, WPawn, BPawn},
			{"", BBishop, "", BPawn, "", WBishop, "", ""},
			{"", BQueen, "", WPawn, "", WKnight, "", ""},
			{"", "", "", "", "", "", "", ""},
			{WPawn, "", "", "", "", WQueen, "", ""},
			{"", WPawn, "", "", "", WPawn, WKing, ""},
			{WRook, "", "", "", "", "", "", WRook},
		},
		AvailableCastles: "-",
		Turn:             "w",
		InPasantSquare:   "-",
		HalfMoves:        1,
		MovesCount:       40,
	}

	FEN := board.FEN()

	assert.Equal(t, "r3r1k1/pp3nPp/1b1p1B2/1q1P1N2/8/P4Q2/1P3PK1/R6R w - - 1 40", FEN)
}

func Test_TranslateFEN(t *testing.T) {
	assert := assert.New(t)
	FEN := "r3r1k1/pp3nPp/1b1p1B2/1q1P1N2/8/P4Q2/1P3PK1/R6R w - - 2 35"
	b := Board{}

	err := b.TranslateFEN(FEN)

	assert.Nil(err, err)
	assert.Equal([][]Piece{
		{BRook, "", "", "", BRook, "", BKing, ""},
		{BPawn, BPawn, "", "", "", BKnight, WPawn, BPawn},
		{"", BBishop, "", BPawn, "", WBishop, "", ""},
		{"", BQueen, "", WPawn, "", WKnight, "", ""},
		{"", "", "", "", "", "", "", ""},
		{WPawn, "", "", "", "", WQueen, "", ""},
		{"", WPawn, "", "", "", WPawn, WKing, ""},
		{WRook, "", "", "", "", "", "", WRook},
	}, b.Board)
	assert.Equal("w", b.Turn)
	assert.Equal("-", b.AvailableCastles)
	assert.Equal("-", b.InPasantSquare)
	assert.Equal(2, b.HalfMoves)
	assert.Equal(35, b.MovesCount)
}

func Test_InitialPosition(t *testing.T) {
	assert := assert.New(t)

	b := NewBoard()

	assert.Equal([][]Piece{
		{BRook, BKnight, BBishop, BQueen, BKing, BBishop, BKnight, BRook},
		{BPawn, BPawn, BPawn, BPawn, BPawn, BPawn, BPawn, BPawn},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{WPawn, WPawn, WPawn, WPawn, WPawn, WPawn, WPawn, WPawn},
		{WRook, WKnight, WBishop, WQueen, WKing, WBishop, WKnight, WRook},
	}, b.Board)
	assert.Equal("w", b.Turn)
	assert.Equal("KQkq", b.AvailableCastles)
	assert.Equal("-", b.InPasantSquare)
	assert.Equal(0, b.HalfMoves)
	assert.Equal(1, b.MovesCount)
	assert.Len(b.MovesHistory, 0)
}
