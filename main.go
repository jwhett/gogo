package main

import (
	"fmt"
)

func main() {
	var board Board
	fmt.Printf("Black: %d\nWhite: %d\n", BLACK, WHITE)
	board.PlayMove(BLACK, "D4")
	board.Show()
}
