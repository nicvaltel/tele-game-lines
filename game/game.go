package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	mb "github.com/nicvaltel/GoUtils/maybe"
)

// type Game = LinesGame[Cell, Color, Input]

func gameLoop(game *GameState) bool {

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
	fmt.Printf("Game Over\n You Win\n") // TODO Implement
}

func log(str string) {
	fmt.Println(str)
}

func getInput() (Input, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter cell in format of e4: ")
	str, _ := reader.ReadString('\n')
	return ParseInput(str)
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
		field:        field_,
		selectedBall: mb.Nothing[Cell]{},
		nextBalls:    nextBalls_,
		prevStates:   []PrevGameState{},
		turnIsEnded:  true,
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
