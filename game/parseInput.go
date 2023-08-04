package game

import (
	"fmt"
	"strings"

	mb "github.com/nicvaltel/GoUtils/maybe"
)

func letterToInt(char string) (Coord, error) {
	char = strings.ToLower(char)
	if len(char) != 1 || char < "a" || char > "z" {
		return 0, fmt.Errorf("invalid input: %s", char)
	}
	return uint8(int(char[0]) - int('a')), nil
}

func numberToInt(char string) (Coord, error) {
	char = strings.ToLower(char)
	if len(char) != 1 || char < "1" || char > "9" {
		return 0, fmt.Errorf("invalid input: %s", char)
	}
	return uint8(int(char[0]) - int('1')), nil
}

func ParseInput(str string) (Input, error) {
	noCell := mb.Nothing[Cell]{}
	if str == "\n" {
		return Input{Comand: INPUT_END_TURN, Cell: noCell}, nil
	}
	if str == "undo\n" {
		return Input{Comand: INPUT_UNDO, Cell: noCell}, nil
	}
	if len(str) == 3 {
		x, err := letterToInt(str[0:1])
		if err != nil {
			return Input{}, err
		}
		y, err := numberToInt(str[1:2])
		if err != nil {
			return Input{}, err
		}
		return Input{Comand: INPUT_COORD, Cell: mb.Just[Cell]{Val: Cell{x, y}}}, nil
	}
	return Input{}, fmt.Errorf("Invalid input: %s", str)
}
