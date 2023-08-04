package game

import (
	"Lines/utils"
	"errors"
	"fmt"
	"math/rand"
	"time"

	mb "github.com/nicvaltel/GoUtils/maybe"
)

const FIELD_SIZE = 9

const COLORS_NUM = 7

const FULL_ROW = 5

const BALLS_PER_TURN = 3

type Field = [][]Color

type Color = uint8 // Zero for the empty cell

type Command uint8

const (
	INPUT_END_TURN = 0
	INPUT_UNDO     = 1
	INPUT_COORD    = 2
)

type Input struct {
	Comand Command
	Cell   mb.Maybe[Cell]
}

type Cell struct {
	Row Coord
	Col Coord
}

type PathField = [][]int

type Coord = uint8

type PrevGameState struct {
	field     Field
	nextBalls []Color
}

type GameState struct {
	field        Field
	selectedBall mb.Maybe[Cell]
	nextBalls    []Color
	turnIsEnded  bool
	pathField    mb.Maybe[PathField]
	prevStates   []PrevGameState
}

func (game *GameState) TurnIsEnded() bool {
	return game.turnIsEnded
}

func undoLastMove(game *GameState) error {
	if len(game.prevStates) == 0 {
		return errors.New("Unable to undo")
	} else {
		prevS := game.prevStates[len(game.prevStates)-1]
		for r, cols := range prevS.field {
			for c, val := range cols {
				game.field[r][c] = val
			}
		}

		for i, val := range prevS.nextBalls {
			game.nextBalls[i] = val
		}

		game.selectedBall = mb.Nothing[Cell]{}
		game.pathField = nil
		game.turnIsEnded = false
		game.prevStates = game.prevStates[:len(game.prevStates)-1]
		return nil
	}
}

func addPrevState(game *GameState) {
	newField := make([][]Color, FIELD_SIZE)
	for i := 0; i < FIELD_SIZE; i++ {
		newField[i] = make([]Color, FIELD_SIZE)
	}

	newNextBalls := make([]Color, BALLS_PER_TURN)

	newState := PrevGameState{field: newField, nextBalls: newNextBalls}

	for r, cols := range game.field {
		for c, val := range cols {
			newState.field[r][c] = val
		}
	}

	for i, val := range game.nextBalls {
		newState.nextBalls[i] = val
	}

	game.prevStates = append(game.prevStates, newState)
}

func (game *GameState) ProcessInput(input Input) error {
	switch input.Comand {
	case INPUT_END_TURN:
		addPrevState(game)
		game.turnIsEnded = true
		return nil
	case INPUT_UNDO:
		err := undoLastMove(game)
		return err
	case INPUT_COORD:
		fmt.Printf("input = %v\n", input)
		cell, _ := input.Cell.FromJust()
		if game.GetColor(cell) == 0 {
			if game.selectedBall.IsJust() {
				if game.PathIsAvailable(cell) {
					game.MoveBall(cell)
					return nil
				} else {
					return errors.New("Path is not available")
				}
			} else {
				return errors.New("Ball is not selected")
			}
		} else {
			err := game.SelectBall(cell)
			if err != nil {
				return err
			}
			return nil
		}
	default:
		return errors.New("Incorrect input")
	}
}

func (game *GameState) PathIsAvailable(cellTo Cell) bool {
	pathField, _ := game.pathField.FromJust()
	return pathField[cellTo.Row][cellTo.Col] >= 0
}

func (game *GameState) GetColor(cell Cell) Color {
	return game.field[cell.Row][cell.Col]
}

func (game *GameState) SelectBall(cell Cell) error {
	if !checkBounds(cell) {
		return errors.New("Out of bounds")
	}
	if game.GetColor(cell) != 0 {
		game.selectedBall = mb.Just[Cell]{Val: cell}
		pathField := FillPathField(game.field, cell)
		game.pathField = mb.Just[PathField]{Val: pathField}
		return nil
	} else {
		return errors.New("Cell is empty")
	}
}

func checkBounds(cell Cell) bool {
	check := func(x Coord) bool {
		return x >= 0 && x < FIELD_SIZE
	}
	return check(cell.Col) && check(cell.Row)
}

func (game *GameState) MoveBall(cell Cell) error {
	if !game.selectedBall.IsJust() {
		return errors.New("No ball selected")
	} else {
		addPrevState(game)
		selBall, _ := game.selectedBall.FromJust()
		color := game.GetColor(selBall)
		game.field[selBall.Row][selBall.Col] = 0
		game.field[cell.Row][cell.Col] = color
		game.ProcessFullLines()
		game.turnIsEnded = true
		return nil
	}
}

func freeCells(field Field) []Cell {
	free := make([]Cell, 0)
	for r, cols := range field {
		for c, color := range cols {
			if color == 0 {
				free = append(free, Cell{Row: uint8(r), Col: uint8(c)})
			}
		}
	}
	return free
}

func (game *GameState) InsertNewBalls() error {
	if !game.TurnIsEnded() {
		return nil
	}

	free := freeCells(game.field)
	if len(free) < BALLS_PER_TURN {
		return errors.New("Not enought free space")
	}

	newRandPositions, err := utils.DifferentRandomNumbers(0, len(free)-1, BALLS_PER_TURN)
	if err != nil {
		return errors.New("Not enought free space")
	}

	for i := 0; i < BALLS_PER_TURN; i++ {
		game.field[free[newRandPositions[i]].Row][free[newRandPositions[i]].Col] = game.nextBalls[i]
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < BALLS_PER_TURN; i++ {
		game.nextBalls[i] = uint8(rand.Intn(COLORS_NUM) + 1)
	}

	game.ProcessFullLines()
	game.turnIsEnded = false
	return nil
}
