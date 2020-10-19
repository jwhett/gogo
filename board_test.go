package gogo_test

import (
	"fmt"
	"testing"

	"github.com/jwhett/gogo"
)

func TestSpace(t *testing.T) {
	var s gogo.Space
	if s.OccupiedBy != 0 {
		t.Fatalf("Space is occupied after initialization. Occupied by %d.", s.OccupiedBy)
	}
}

func TestSpace_IsOccupied(t *testing.T) {
	var s gogo.Space
	if s.IsOccupied() {
		t.Fatal("Space didn't initialize to zero")
	}
}

func TestSpace_Occupy(t *testing.T) {
	var s gogo.Space
	s.Occupy(gogo.BLACK)
	if !s.IsOccupied() {
		t.Fatal("Failed to occupy")
	}
	t.Logf("Occupied by: %d", s)
}

func TestRow(t *testing.T) {
	var r gogo.Row
	if len(r.Spaces) != 19 {
		t.Fatalf("Row doesn't have 19 spaces: %d", len(r.Spaces))
	}
}

func TestBoard(t *testing.T) {
	var b gogo.Board
	if len(b.Rows) != 19 {
		t.Fatalf("Board doesn't have 19 Rows: %d", len(b.Rows))
	}
}

func TestBoard_GetSpace(t *testing.T) {
	var b gogo.Board
	s, err := b.GetSpace("A1")
	if err != nil {
		t.Fatalf("Failed to get space at A1")
	}
	if fmt.Sprintf("%T", s) != "*gogo.Space" {
		t.Fatalf("Failed to return a Space - got %T instead", s)
	}
}

func TestBoard_PlayMove(t *testing.T) {
	var b gogo.Board
	coords := "A1"
	moveErr := b.PlayMove(gogo.BLACK, coords)
	if moveErr != nil {
		t.Fatalf("Failed to play move at %s. Error: %v", coords, moveErr)
	}
	space, _ := b.GetSpace(coords)
	if !space.IsOccupied() {
		t.Fatal("Didn't receive error when playing, but space is not occupied...")
	}
}
