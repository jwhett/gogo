package gogo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jwhett/gogo/sgf"
)

const (
	xletters string = " abcdefghijklmnopqrs"
	yletters string = " srqponmlkjihgfedcba"
)

// Space defines spaces on a Go board.
type Space struct {
	OccupiedBy Player
}

// Row is an array of 19 Spaces.
type Row struct {
	Spaces [19]Space
}

// Board is an array of 19 Rows to reflect a full
// Go board.
type Board struct {
	Rows    [19]Row
	History string
}

// IsOccupied returns true if no Players occupy the space.
func (s *Space) IsOccupied() bool {
	return s.OccupiedBy > 0
}

// Occupy will fill an unoccupied space with a
// Player's token. Returns an error if one attempts
// to play in an occupied Space.
func (s *Space) Occupy(p Player) error {
	if s.IsOccupied() {
		return fmt.Errorf("space is taken")
	}
	s.OccupiedBy = p
	return nil
}

func (r Row) String() string {
	spaces := make([]string, 19)
	for i, s := range r.Spaces {
		spaces[i] = s.String()
	}
	return strings.Join(spaces, " ")
}

func (s Space) String() string {
	if s.IsOccupied() {
		return fmt.Sprint(s.OccupiedBy)
	}
	return "-"
}

// Show will print the state of the Board.
func (b *Board) Show() {
	fmt.Println("   A B C D E F G H I J K L M N O P Q R S")
	for i := 18; i >= 0; i-- {
		fmt.Printf("%2d|%v|%2d\n", i+1, b.Rows[i], i+1)
	}
	fmt.Println("   A B C D E F G H I J K L M N O P Q R S")
}

// GetSpace attempts to return a Space from a given
// set of coordinates.
func (b *Board) GetSpace(coords string) (*Space, error) {
	y, x, parseErr := ParseCoordinates(coords)
	if parseErr != nil {
		return nil, parseErr
	}
	return &b.Rows[x-1].Spaces[y-1], nil
}

// PlayMove will attempt to occupy a space
// with a Player's token from the Board level.
func (b *Board) PlayMove(p Player, m string) error {
	s, err := b.GetSpace(m)
	if err != nil {
		return err
	}
	occErr := s.Occupy(p)
	if occErr != nil {
		return occErr
	}
	// We can't error here if b.GetSpace() succeeds.
	x, y, _ := ParseCoordinates(m)
	b.History += sgf.EmitMove(int(p), x, y)
	return nil
}

// ShowHistory will print the game history in SGF
// format through the current move.
func (b *Board) ShowHistory() {
	fmt.Printf("%s\n%s)", sgf.Header(), b.History)
}

// ParseCoordinates is a helper function to parse
// NX[|X] coordinates.
func ParseCoordinates(s string) (int, int, error) {
	// split NXX move into X,Y
	// e.g. a1 = 1,1
	sx := s[:1]
	x := strings.Index(xletters, sx)
	y, yerr := strconv.Atoi(s[1:])
	if yerr != nil {
		return 0, 0, yerr
	}
	if (x < 0) || (x > 19) {
		return 0, 0, fmt.Errorf("X is not a possible coordinate: %d", x)
	}
	if (y < 0) || (y > 19) {
		return 0, 0, fmt.Errorf("Y is not a possible coordinate: %d", y)
	}
	return x, y, nil
}
