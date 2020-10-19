package main

import (
	"fmt"
)

type Player int

const (
	BLACK Player = iota + 1
	WHITE
)

type Space struct {
	occupiedBy Player
}

type Row struct {
	Spaces [19]Space
}

type Board struct {
	Rows [19]Row
}

func (s Space) IsOccupied() bool {
    return s.occupiedBy > 0
}

func (s Space) String() string {
	if s.IsOccupied() {
		return string(s.occupiedBy)
	} else {
		return "0"
	}
}

func (b Board) Show() {
	for i := range b.Rows {
		fmt.Println(b.Rows[i])
	}
}

