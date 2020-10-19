package main

import (
	"fmt"
)

// Space defines spaces on a Go board.
type Space struct {
	occupiedBy Player
}

// Row is an array of 19 Spaces.
type Row struct {
	Spaces [19]Space
}

// Board is an array of 19 Rows to reflect a full
// Go board.
type Board struct {
	Rows [19]Row
}

// IsOccupied returns true if no Players occupy the space.
func (s Space) IsOccupied() bool {
	return s.occupiedBy > 0
}

func (s Space) String() string {
	if s.IsOccupied() {
		return fmt.Sprint(s.occupiedBy)
	}
	return "0"
}

// Show will print the state of the Board.
func (b Board) Show() {
	for i := range b.Rows {
		fmt.Println(b.Rows[i])
	}
}
