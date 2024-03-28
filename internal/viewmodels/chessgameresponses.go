package viewmodels

type (
	ChessGameResponse struct {
		MoveDone       bool     `json:"move_done"`
		IsCheckMate    bool     `json:"is_checkmate"`
		IsStaleMate    bool     `json:"is_stalemate"`
		AvailableMoves []string `json:"available_moves"`
		FEN            string   `json:"fen"`
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)
