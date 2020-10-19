package main

import (
	"fmt"

	gb "github.com/jwhett/gogo"
)

func main() {
	var board gb.Board
	fmt.Printf("Black: %d\nWhite: %d\n", gb.BLACK, gb.WHITE)
	board.PlayMove(gb.BLACK, "A1")
	board.PlayMove(gb.WHITE, "S19")
	board.Show()
}
