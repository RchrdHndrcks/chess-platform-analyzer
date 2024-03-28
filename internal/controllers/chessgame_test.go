package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"chenizz/internal/controllers/internal"
	"chenizz/internal/services/mocks"
	"chenizz/internal/viewmodels"
)

func Test_MakeMove_ValidMove(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	serviceResponse := viewmodels.ChessGameResponse{
		MoveDone:       true,
		IsCheckMate:    false,
		IsStaleMate:    false,
		AvailableMoves: []string{"e7e6"},
		FEN:            "rnbqkbnr/pppp1ppp/4p3/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1",
	}

	makeMoveServiceMock := mocks.ChessGameServiceMock{}
	makeMoveServiceMock.PatchMakeMove(serviceResponse, nil)

	controller := ChessGameController{
		makeMoveServiceMock,
	}

	reqParams := internal.MakeMoveParams{
		Move: "e2e4",
		FEN:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
	}
	body, _ := json.Marshal(reqParams)

	req, err := http.NewRequest("POST", "/game/chess/make-move", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("error calling http.NewRequest: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.MakeMove)

	// Act
	handler.ServeHTTP(rr, req)
	resp := viewmodels.ChessGameResponse{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	// Assert
	assert.Equal(http.StatusOK, rr.Code)
	assert.Equal(serviceResponse, resp)
}

func Test_MakeMove_NotValidMove(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	err := fmt.Errorf("error in service")
	serviceResponse := viewmodels.ChessGameResponse{
		MoveDone:       false,
		IsCheckMate:    false,
		IsStaleMate:    false,
		AvailableMoves: []string{"d1b2"},
		FEN:            "8/8/8/8/1BP5/1BKP4/1b1P4/1k1NR3 w - - 0 1",
	}
	serviceMock := mocks.ChessGameServiceMock{}
	serviceMock.PatchMakeMove(serviceResponse, err)

	controller := ChessGameController{serviceMock}

	reqParams := internal.MakeMoveParams{
		Move: "a1a2",
		FEN:  "8/8/8/8/1BP5/1BKP4/1b1P4/1k1NR3 w - - 0 1",
	}
	body, _ := json.Marshal(reqParams)

	req, err := http.NewRequest("POST", "/game/chess/make-move", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("error calling http.NewRequest: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.MakeMove)

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(http.StatusUnprocessableEntity, rr.Code)
	assert.Equal("{\"error\":\"error in service\"}\n", rr.Body.String())
}
