// Package sgf describes the SGF file formatting and provides functions
// for working with SGF data. SGF FF[4] is the current file format
// as of this writing. See https://www.red-bean.com/sgf/ and
// https://www.red-bean.com/sgf/proplist_t.html
package sgf

import (
	"fmt"
)

const (
	// xletters is a string to translate numbers
	// into letter coordinates.
	xletters string = " abcdefghijklmnopqrs"
	// yletters is a string to translate numbers
	// into letter coordinates.
	yletters = " srqponmlkjihgfedcba"
	// Application describes the application used
	// to generate the SGF file.
	Application = "AP[%s]"
	// BlackPass shows that the Black player passed.
	BlackPass = "B[]"
	// BlackMove describes a move with xy coordinates
	// a-s.
	BlackMove = "B[%2s]"
	// BlackRank describes the player's rank.
	BlackRank = "BR[%s]"
	// Comment is a free-form comment string.
	Comment = "C[%s]"
	// DateTime is the date and time of the game. ISO date - YYYY-MM-DD.
	DateTime = "DT[%s]"
	// FileFormat describes the SGF version.
	FileFormat = "FF[%d]"
	// Game is the type of game being represented. GM[1] is Go/Baduk.
	Game = "GM[%d]"
	// GameName is the name of the game record.
	GameName = "GN[%s]"
	// Komi given to the Black player. This is either 0 or
	// must have a half of a point to break ties. 6.5 is a
	// common Komi granted to Black.
	Komi = "KM[%.1f]"
	// PlayerBlack is the Black player's name.
	PlayerBlack = "PB[%s]"
	// Place is where the game is being played. Geographically,
	// application, web site, etc...
	Place = "PC[%s]"
	// PlayerWhite is the White player's name.
	PlayerWhite = "PW[%s]"
	// Result describes the winner of the game and by
	// what margin. (e.g. B+49.5 => Black wins by 49.5 points.
	// W+Resign => White wins by resignation)
	Result = "RE[%s]"
	// Rules describes the counting rules.
	Rules = "RU[%s]"
	// Source is the source of the SGF file.
	Source = "SO[%s]"
	// Size is the size of the Go board. Standard, square sizes
	// are 9, 13, 19.
	Size = "SZ[%d]"
	// WhiteMove describes a move with xy coordinates
	// a-s.
	WhiteMove = "W[%2s]"
	// WhitePass shows that the White player passed.
	WhitePass = "W[]"
	// WhiteRank is the White player's rank.
	WhiteRank = "WR[%s]"
)

// Header builds the header of the SGF output.
func Header() (header string) {
	header = "(;"
	header += fmt.Sprintf(FileFormat, 4)
	header += fmt.Sprintf(Game, 1)
	header += fmt.Sprintf(Size, 19)
	header += fmt.Sprintf(Application, "GoGo")
	return
}

// EmitMove will emit an SGF-formatted move
func EmitMove(playerID, x, y int) string {
	var move string
	switch playerID {
	case 1:
		move = BlackMove
	case 2:
		move = WhiteMove
	}
	translated := string(xletters[x]) + string(yletters[y])
	return fmt.Sprintf(";"+move, translated)
}
