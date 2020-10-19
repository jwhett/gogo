package gogo_test

import (
	"testing"

	"github.com/jwhett/gogo"
)

func TestSpace_Init(t *testing.T) {
	var s gogo.Space
	if s.IsOccupied() {
		t.Fatal("Space didn't initialize to zero")
	}
}
