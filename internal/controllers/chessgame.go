package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"chenizz/internal/controllers/internal"
	"chenizz/internal/interfaces"
	"chenizz/internal/viewmodels"
)

type ChessGameController struct {
	interfaces.IChessGameService
}

func (c ChessGameController) MakeMove(w http.ResponseWriter, r *http.Request) {
	p, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse(err))
		return
	}

	params := internal.MakeMoveParams{}
	err = json.Unmarshal(p, &params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse(err))
		return
	}

	resp, err := c.IChessGameService.MakeMove(params.Move, params.FEN)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(errorResponse(err))
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func errorResponse(err error) viewmodels.ErrorResponse {
	return viewmodels.ErrorResponse{Error: err.Error()}
}
