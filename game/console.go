package game

import "fmt"

func (game *GameState) DysplayScreen() error {

	fmt.Printf("GameStateA: %+v\n\n", game)

	fmt.Printf("Next Balls: %v\n", game.nextBalls)
	fmt.Printf("Selected Ball: %v\n", game.selectedBall)

	arr := [9]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	fmt.Println(" | 1 2 3 4 5 6 7 8 9")
	fmt.Println("____________________")
	for i, cols := range game.field {
		fmt.Print(arr[i] + "|")
		for _, color := range cols {
			fmt.Printf(" %v", color)
		}
		fmt.Println()
	}
	fmt.Printf("---------------------------------------------------\n\n")
	return nil
}
