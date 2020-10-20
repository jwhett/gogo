package main

import (
	"fmt"

    "github.com/jwhett/gogo"
	"github.com/jwhett/gogo/sgf"
)

func main() {
	var board gogo.Board
	fmt.Printf("Black: %d\nWhite: %d\n", gogo.BLACK, gogo.WHITE)
	board.PlayMove(gogo.BLACK, "A1")
	board.PlayMove(gogo.WHITE, "S19")
	board.Show()
	fmt.Printf(sgf.Application, "GoGo")
}
