package main

import (
	"fmt"

	"github.com/jwhett/gogo"
)

func main() {
	var board gogo.Board
	fmt.Printf("Black: %d\nWhite: %d\n", gogo.BLACK, gogo.WHITE)
	board.PlayMove(gogo.BLACK, "d4")
	board.PlayMove(gogo.WHITE, "f3")
	board.PlayMove(gogo.BLACK, "c6")
	board.PlayMove(gogo.WHITE, "a19")
	board.PlayMove(gogo.BLACK, "s1")
	board.Show()
	board.ShowHistory()
}
