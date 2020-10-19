package gogo

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRS"
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
func (s *Space) IsOccupied() bool {
	return s.occupiedBy > 0
}

// Occupy will fill an unoccupied space with a
// Player's token. Returns an error if one attempts
// to play in an occupied Space.
func (s *Space) Occupy(p Player) error {
	if s.IsOccupied() {
		return fmt.Errorf("space is taken")
	}
	s.occupiedBy = p
	return nil
}

func (s Space) String() string {
	if s.IsOccupied() {
		return fmt.Sprint(s.occupiedBy)
	}
	return "0"
}

// Show will print the state of the Board.
func (b *Board) Show() {
	for i := range b.Rows {
		fmt.Println(b.Rows[i])
	}
}

// GetSpace attempts to return a Space from a given
// set of coordinates.
func (b *Board) GetSpace(coords string) (*Space, error) {
	x, y, parseErr := parseCoordinates(coords)
	if parseErr != nil {
		return nil, parseErr
	}
	row := b.Rows[x]
	space := row.Spaces[y]
	return &space, nil
}

// PlayMove will attempt to occupy a space
// with a Player's token from the Board level.
func (b *Board) PlayMove(p Player, m string) error {
	s, err := b.GetSpace(m)
	if err != nil {
		return err
	}
	s.Occupy(p)
	return nil
}

func parseCoordinates(s string) (x int, y int, err error) {
	// split NXX move into X,Y
	x = strings.Index(letters, s[:1]) + 1
	y, err = strconv.Atoi(s[1:])
	return
}
