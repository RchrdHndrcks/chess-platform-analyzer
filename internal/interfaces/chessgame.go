package interfaces

import "chenizz/internal/viewmodels"

type IChessGameService interface {
	MakeMove(move string, fen string) (viewmodels.ChessGameResponse, error)
}
