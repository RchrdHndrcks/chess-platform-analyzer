package mocks

import "chenizz/internal/viewmodels"

type ChessGameServiceMock struct {
	response viewmodels.ChessGameResponse
	err      error
}

func (c *ChessGameServiceMock) PatchMakeMove(r viewmodels.ChessGameResponse, err error) {
	c.response = r
	c.err = err
}

func (c ChessGameServiceMock) MakeMove(move string, fen string) (viewmodels.ChessGameResponse, error) {
	return c.response, c.err
}
