package game

import mb "github.com/nicvaltel/GoUtils/maybe"

type LinesGame[CELL, COLOR, INPUT any] interface {
	SelectBall(CELL) error
	MoveBall(CELL) error
	FindFullLines() [][]CELL
	InsertNewBalls() error
	CheckWinCondition() bool
	CheckLooseCondition() bool
	GetColor(CELL) COLOR
	UpdateGame(CELL)
	ProcessInput(INPUT) error
	PathIsAvailable(CELL) mb.Maybe[[]CELL]
	TurnIsEnded() bool
}

type LinesVisual[CELL, COLOR any] interface {
	DysplayScreen() error
}
