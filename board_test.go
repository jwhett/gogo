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

func TestSpace_OccupyFailure(t *testing.T) {
	var s gogo.Space
	s.Occupy(gogo.BLACK)
	err := s.Occupy(gogo.WHITE)
	if err == nil {
		t.Fatal("No error when trying to play in an occupied Space.")
	}
}

func TestRow(t *testing.T) {
	var r gogo.Row
	if len(r.Spaces) != 19 {
		t.Fatalf("Row doesn't have 19 spaces: %d", len(r.Spaces))
	}
}

func TestParseCoordinates_Simple(t *testing.T) {
	x, y, err := gogo.ParseCoordinates("a1")
	if err != nil {
		t.Fatalf("Error while parsing a1: %v", err)
	}
	if x != 1 {
		t.Fatalf("Failed to map A to 1. Got: %d", x)
	}
	if y != 1 {
		t.Fatalf("Failed to properly set y coord to 1. Got: %d", y)
	}
}

func TestParseCoordinates_Complex(t *testing.T) {
	x, y, err := gogo.ParseCoordinates("s19")
	if err != nil {
		t.Fatalf("Error while parsing s19: %v", err)
	}
	if x != 19 {
		t.Fatalf("Failed to map s to 19. Got: %d", x)
	}
	if y != 19 {
		t.Fatalf("Failed to properly set y coord to 19. Got: %d", y)
	}
}

func TestParseCoordinates_OOB(t *testing.T) {
	_, _, err := gogo.ParseCoordinates("t20")
	if err == nil {
		t.Fatal("Didn't fail when out of bounds")
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
	s, err := b.GetSpace("a1")
	if err != nil {
		t.Fatalf("Failed to get space at a1")
	}
	if fmt.Sprintf("%T", s) != "*gogo.Space" {
		t.Fatalf("Failed to return a Space - got %T instead", s)
	}
}

func TestBoard_PlayMove(t *testing.T) {
	var b gogo.Board
	coords := "a1"
	moveErr := b.PlayMove(gogo.BLACK, coords)
	if moveErr != nil {
		t.Fatalf("Failed to play move at %s. Error: %v", coords, moveErr)
	}
	space, _ := b.GetSpace(coords)
	if !space.IsOccupied() {
		t.Fatal("Didn't receive error when playing, but space is not occupied...")
	}
}
