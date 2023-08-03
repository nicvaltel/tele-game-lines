package game

import "fmt"

func (game *GameState) DysplayScreen() error {

	fmt.Printf("Next Balls: %v\n", game.nextBalls)
	fmt.Printf("Selected Ball: %v\n", game.selectedBall)

	arr := [9]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	fmt.Println(" |123456789")
	fmt.Println("___________")
	for i, cols := range game.field {
		fmt.Print(arr[i] + "|")
		for _, color := range cols {
			fmt.Print(color)
		}
		fmt.Println()
	}
	fmt.Printf("---------------------------------------------------\n\n")
	return nil
}
