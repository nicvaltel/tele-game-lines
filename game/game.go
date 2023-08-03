package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	mb "github.com/nicvaltel/GoUtils/maybe"
)

type Game = LinesGame[Cell, Color, Input]

func gameLoop(game *GameState) bool {

	fmt.Printf("GameStateA: %+v\n\n", game)

	err := game.InsertNewBalls()
	if err != nil {
		loseGame(*game)
		return false
	}

	game.DysplayScreen()

	input, err := getInput()
	if err != nil {
		log(err.Error())
		return false
	}

	err = game.ProcessInput(input)
	if err != nil {
		log(err.Error())
		return true
	}

	return true
}

func loseGame(game GameState) {
	fmt.Printf("Game Over\n You Lose\n") // TODO Implement
}

func winGame(game GameState) {
	panic("unimplemented")
}

func log(str string) {
	fmt.Println(str)
}

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

func parseInput(str string) (Input, error) {
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

func getInput() (Input, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter cell in format of e4: ")
	str, _ := reader.ReadString('\n')
	res, err := parseInput(str)
	if err != nil {
		return res, nil
	}
	return parseInput(str)
}

func initialGameState() GameState {

	field_ := make([][]Color, FIELD_SIZE)
	for i := range field_ {
		field_[i] = make([]Color, FIELD_SIZE)
	}

	nextBalls_ := make([]Color, BALLS_PER_TURN)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < BALLS_PER_TURN; i++ {
		nextBalls_[i] = uint8(rand.Intn(COLORS_NUM) + 1)
	}

	game := GameState{
		field:         field_,
		selectedBall:  mb.Nothing[Cell]{},
		nextBalls:     nextBalls_,
		prevField:     field_,
		prevNextBalls: nextBalls_,
		turnIsEnded:   true,
	}
	return game
}

func RunGame() {
	fmt.Println("OKAY")

	game := initialGameState()

	contininueGame := true
	for contininueGame {
		contininueGame = gameLoop(&game)
	}
}
