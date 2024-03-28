package services

import (
	"fmt"

	"chenizz/internal/services/internal/chess"
	"chenizz/internal/viewmodels"
)

type ChessGameService struct {
}

func (c ChessGameService) MakeMove(move string, fen string) (viewmodels.ChessGameResponse, error) {
	b := chess.NewBoard()
	err := b.TranslateFEN(fen)
	if err != nil {
		return viewmodels.ChessGameResponse{}, fmt.Errorf("error calling b.TranslateFEN: %w", err)
	}

	availableMoves := b.AvailableLegalMoves()
	if !contains(availableMoves, move) {
		return viewmodels.ChessGameResponse{AvailableMoves: availableMoves},
			fmt.Errorf("%s is not a legal move", move)
	}

	b.MakeMove(move)
	availableMoves = b.AvailableLegalMoves()

	response := viewmodels.ChessGameResponse{
		MoveDone:       true,
		AvailableMoves: availableMoves,
		FEN:            b.FEN(),
	}

	response.IsCheckMate = availableMoves == nil
	response.IsStaleMate = availableMoves != nil && len(availableMoves) == 0
	return response, nil
}

func contains(s []string, m string) bool {
	for _, v := range s {
		if v == m {
			return true
		}
	}

	return false
}
