package main

import "chenizz/internal/controllers"

type IServiceContainer interface {
	ChessGameController() controllers.ChessGameController
}

type k struct{}

func (k k) ChessGameController() controllers.ChessGameController {
	return controllers.ChessGameController{}
}

func ServiceContainer() IServiceContainer {
	return k{}
}
