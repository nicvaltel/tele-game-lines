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

type Coord = uint8

type GameState struct {
	field         Field
	selectedBall  mb.Maybe[Cell]
	nextBalls     []Color
	prevField     Field
	prevNextBalls []Color
	turnIsEnded   bool
}

func (game *GameState) TurnIsEnded() bool {
	return game.turnIsEnded
}

func undoLastMove(game *GameState) error {
	if game.prevField == nil {
		return errors.New("Unable to undo")
	}
	game.field = game.prevField
	game.prevField = nil
	game.nextBalls = game.prevNextBalls
	game.prevNextBalls = nil
	game.turnIsEnded = true
	return nil
}

func (game *GameState) ProcessInput(input Input) error {
	switch input.Comand {
	case INPUT_END_TURN:
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
				if game.PathIsAvailable(cell).IsJust() {
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

// func neighbourCells(game GameState, cell Cell) []Cell {

// 	neighbours := make([]Cell, 0)

// 	appendNeighbour := func(cl Cell) {
// 		if checkBounds(cl) && game.GetColor(cl) == 0 {
// 			neighbours = append(neighbours, cl)
// 		}
// 	}

// 	r := cell.row
// 	c := cell.col

// 	appendNeighbour(Cell{r + 1, c})
// 	appendNeighbour(Cell{r - 1, c})
// 	appendNeighbour(Cell{r, c + 1})
// 	appendNeighbour(Cell{r, c - 1})

// 	return neighbours
// }

func findPath(game GameState, cellTo Cell, path []Cell) mb.Maybe[[]Cell] {
	panic("undefined")
}

// func findPath(game GameState, cellTo Cell, path []Cell) mb.Maybe[[]Cell] {
// 	if path[len(path)-1] == cellTo {
// 		return mb.Nothing[[]Cell]{}
// 	} else {
// 		neighbours := neighbourCells(game, path[len(path)-1])
// 		if len(neighbours) == 0 {
// 			return mb.Nothing[[]Cell]{}
// 		} else {
// 			// candidates := make([][]Cell,0)
// 			for _, nb := range neighbours {
// 				path = append(path, nb)
// 				res := findPath(game, cellTo, path)
// 				if res.IsJust() {
// 					return res
// 				}
// 			}
// 		}
// 	}
// 	return mb.Nothing[[]Cell]{}
// }

func (game *GameState) PathIsAvailable(cellTo Cell) mb.Maybe[[]Cell] {
	cellFrom, _ := game.selectedBall.FromJust()
	path := []Cell{cellFrom}
	return findPath(*game, cellTo, path)
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

func (game *GameState) FindFullLines() [][]Cell {
	panic("undefined")
}

func (game *GameState) MoveBall(cell Cell) error {
	game.turnIsEnded = true
	panic("undefined")
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

	game.turnIsEnded = false
	return nil
}
