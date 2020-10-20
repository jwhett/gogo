package sgf_test

import (
    	"fmt"
	"testing"
	"time"

	"github.com/jwhett/gogo/sgf"
)

// Test application format string
func TestFormatting_AP(t *testing.T) {
    if fmt.Sprintf(sgf.Application, "GoGo") != "AP[GoGo]" {
        t.Fatal("Formatting failed for: Application")
    }
}

// Test BlackRank
func TestFormatting_BR(t *testing.T) {
    if fmt.Sprintf(sgf.BlackRank, "9d") != "BR[9d]" {
        t.Fatal("Formatting failed for: BlackRank")
    }
}

// Test Comment
func TestFormatting_C(t *testing.T) {
    if fmt.Sprintf(sgf.Comment, "Testing") != "C[Testing]" {
        t.Fatal("Formatting failed for: Comment")
    }
}

// Test DateTime
func TestFormatting_DT(t *testing.T) {
    testDate := time.Now()
    if fmt.Sprintf(sgf.DateTime, testDate.Local()) != fmt.Sprintf("DT[%s]", testDate.Local()) {
        t.Fatal("Formatting failed for: DateTime")
    }
}

// Test FileFormat
func TestFormatting_FF(t *testing.T) {
    if fmt.Sprintf(sgf.FileFormat, 4) != "FF[4]" {
        t.Fatal("Formatting failed for: FileFormat")
    }
}

// Test Game
func TestFormatting_GM(t *testing.T) {
    if sgf.Game != "GM[1]" {
        t.Fatal("Formatting failed for: Game")
    }
}

// Test GameName
func TestFormatting_GN(t *testing.T) {
    if fmt.Sprintf(sgf.GameName, "Testing") != "GN[Testing]" {
        t.Fatal("Formatting failed for: GameName")
    }
}

// Test Komi
func TestFormatting_KM(t *testing.T) {
    if fmt.Sprintf(sgf.Komi, 6.55) != "KM[6.5]" {
        t.Fatal("Formatting failed for: Komi")
    }
}

// Test PlayerBlack
func TestFormatting_PB(t *testing.T) {
    if fmt.Sprintf(sgf.PlayerBlack, "Testing") != "PB[Testing]" {
        t.Fatal("Formatting failed for: PlayerBlack")
    }
}

// Test Place
func TestFormatting_PC(t *testing.T) {
    if fmt.Sprintf(sgf.Place, "GoGo") != "PC[GoGo]" {
        t.Fatal("Formatting failed for: Place")
    }
}

// Test PlayerWhite
func TestFormatting_PW(t *testing.T) {
    if fmt.Sprintf(sgf.PlayerWhite, "Testing") != "PW[Testing]" {
        t.Fatal("Formatting failed for: PlayerWhite")
    }
}

// Test Result
func TestFormatting_RE(t *testing.T) {
    if fmt.Sprintf(sgf.Result, "B+49.5") != "RE[B+49.5]" {
        t.Fatal("Formatting failed for: Result")
    }
}

// Test Rules
func TestFormatting_RU(t *testing.T) {
    if fmt.Sprintf(sgf.Rules, "Japanese") != "RU[Japanese]" {
        t.Fatal("Formatting failed for: Rules")
    }
}

// Test Source
func TestFormatting_SO(t *testing.T) {
    if fmt.Sprintf(sgf.Source, "Testing") != "SO[Testing]" {
        t.Fatal("Formatting failed for: Source")
    }
}

// Test Size
func TestFormatting_SZ(t *testing.T) {
    if fmt.Sprintf(sgf.Size, 19) != "SZ[19]" {
        t.Fatal("Formatting failed for: Size")
    }
}

// Test WhiteRank
func TestFormatting_WR(t *testing.T) {
    if fmt.Sprintf(sgf.WhiteRank, "9d") != "WR[9d]" {
        t.Fatal("Formatting failed for: WhiteRank")
    }
}
