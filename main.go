package main

import (
	"fmt"
)

func main() {
	var board Board
	fmt.Printf("Black: %d\nWhite: %d\n", BLACK, WHITE)
	board.Show()
}

/*
	letters := [19]string{"A", "B", "C", "D",
                              "E", "F", "G", "H",
                              "I", "J", "K", "L",
			      "M", "N", "O", "P",
                              "Q", "R", "S"}
*/
