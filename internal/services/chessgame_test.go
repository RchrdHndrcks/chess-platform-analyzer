package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeMove_MoveIsValid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	providedFen := "8/5pkp/6p1/8/8/6K1/3P1PPP/8 w - - 0 1"
	providedMove := "d2d4"
	availableMovesForBlack := []string{"f7f6", "f7f5", "g6g5", "h7h6", "h7h5",
		"g7f6", "g7h6", "g7f8", "g7g8", "g7h8"}

	c := ChessGameService{}

	// Act
	r, err := c.MakeMove(providedMove, providedFen)

	// Assert
	assert.Nil(err)
	assert.True(r.MoveDone)
	assert.False(r.IsCheckMate)
	assert.False(r.IsStaleMate)
	assert.ElementsMatch(availableMovesForBlack, r.AvailableMoves)
}

func Test_MakeMove_MoveIsNotValid(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	providedFen := "8/5pkp/6p1/8/8/6K1/3P1PPP/8 w - - 0 1"
	providedMove := "a2a3"
	availableMovesForWhite := []string{"d2d3", "d2d4", "f2f3", "f2f4", "h2h3", "h2h4",
		"g3f3", "g3f4", "g3g4", "g3h3", "g3h4"}

	c := ChessGameService{}

	// Act
	r, err := c.MakeMove(providedMove, providedFen)

	// Assert
	assert.EqualError(err, "a2a3 is not a legal move")
	assert.False(r.MoveDone)
	assert.False(r.IsCheckMate)
	assert.False(r.IsStaleMate)
	assert.ElementsMatch(availableMovesForWhite, r.AvailableMoves)
}

func Test_MakeMove_InvalidFEN(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	providedFen := "invalid-fen"
	providedMove := "a2a3"

	c := ChessGameService{}

	// Act
	r, err := c.MakeMove(providedMove, providedFen)

	// Assert
	assert.EqualError(err, "error calling b.TranslateFEN: invalid FEN")
	assert.False(r.MoveDone)
	assert.False(r.IsCheckMate)
	assert.False(r.IsStaleMate)
	assert.Nil(r.AvailableMoves)
}

func Test_MakeMove_MoveIsCheckMate(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	providedFen := "8/8/8/8/1BP5/1BKP4/1b1P4/1k1NR3 w - - 0 1"
	providedMove := "d1b2"

	c := ChessGameService{}

	// Act
	r, err := c.MakeMove(providedMove, providedFen)

	// Assert
	assert.Nil(err)
	assert.True(r.MoveDone)
	assert.True(r.IsCheckMate)
	assert.False(r.IsStaleMate)
	assert.Nil(r.AvailableMoves)
}

func Test_MakeMove_MoveIsStaleMate(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	providedFen := "k7/2K5/8/8/3Q4/8/8/8 w - - 0 1"
	providedMove := "d4b6"

	c := ChessGameService{}

	// Act
	r, err := c.MakeMove(providedMove, providedFen)

	// Assert
	assert.Nil(err)
	assert.True(r.MoveDone)
	assert.False(r.IsCheckMate)
	assert.True(r.IsStaleMate)
	assert.Empty(r.AvailableMoves)
}
