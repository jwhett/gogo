package main

import (
	"fmt"

	gb "github.com/jwhett/gogo"
)

func main() {
	var board gb.board.Board
	fmt.Printf("Black: %d\nWhite: %d\n", BLACK, WHITE)
	board.PlayMove(BLACK, "A1")
	board.Show()
}
