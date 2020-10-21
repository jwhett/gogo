package sgf_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jwhett/gogo/sgf"
)

func TestFormatting_AP(t *testing.T) {
	if fmt.Sprintf(sgf.Application, "GoGo") != "AP[GoGo]" {
		t.Fatal("Formatting failed for: Application")
	}
	t.Logf("Application: %s", fmt.Sprintf(sgf.Application, "GoGo"))
}

func TestFormatting_BP(t *testing.T) {
	if sgf.BlackPass != "B[]" {
		t.Fatal("Formatting failed for: BlackPass")
	}
	t.Logf("BlackPass: %s", sgf.BlackPass)
}

func TestFormatting_BR(t *testing.T) {
	if fmt.Sprintf(sgf.BlackRank, "9d") != "BR[9d]" {
		t.Fatal("Formatting failed for: BlackRank")
	}
	t.Logf("BlackRank: %s", fmt.Sprintf(sgf.BlackRank, "9d"))
}

func TestFormatting_C(t *testing.T) {
	if fmt.Sprintf(sgf.Comment, "Testing") != "C[Testing]" {
		t.Fatal("Formatting failed for: Comment")
	}
}

func TestFormatting_DT(t *testing.T) {
	testDate := time.Now().Format("2006-01-02")
	if fmt.Sprintf(sgf.DateTime, testDate) != fmt.Sprintf("DT[%s]", testDate) {
		t.Fatalf("Formatting failed for: DateTime.")
	}
}

func TestFormatting_FF(t *testing.T) {
	if fmt.Sprintf(sgf.FileFormat, 4) != "FF[4]" {
		t.Fatal("Formatting failed for: FileFormat")
	}
}

func TestFormatting_GM(t *testing.T) {
	if fmt.Sprintf(sgf.Game, 1) != "GM[1]" {
		t.Fatal("Formatting failed for: Game")
	}
}

func TestFormatting_GN(t *testing.T) {
	if fmt.Sprintf(sgf.GameName, "Testing") != "GN[Testing]" {
		t.Fatal("Formatting failed for: GameName")
	}
}

func TestFormatting_KM(t *testing.T) {
	if fmt.Sprintf(sgf.Komi, 6.55) != "KM[6.5]" {
		t.Fatal("Formatting failed for: Komi")
	}
}

func TestFormatting_PB(t *testing.T) {
	if fmt.Sprintf(sgf.PlayerBlack, "Testing") != "PB[Testing]" {
		t.Fatal("Formatting failed for: PlayerBlack")
	}
}

func TestFormatting_PC(t *testing.T) {
	if fmt.Sprintf(sgf.Place, "GoGo") != "PC[GoGo]" {
		t.Fatal("Formatting failed for: Place")
	}
}

func TestFormatting_PW(t *testing.T) {
	if fmt.Sprintf(sgf.PlayerWhite, "Testing") != "PW[Testing]" {
		t.Fatal("Formatting failed for: PlayerWhite")
	}
}

func TestFormatting_RE(t *testing.T) {
	if fmt.Sprintf(sgf.Result, "B+49.5") != "RE[B+49.5]" {
		t.Fatal("Formatting failed for: Result")
	}
}

func TestFormatting_RU(t *testing.T) {
	if fmt.Sprintf(sgf.Rules, "Japanese") != "RU[Japanese]" {
		t.Fatal("Formatting failed for: Rules")
	}
}

func TestFormatting_SO(t *testing.T) {
	if fmt.Sprintf(sgf.Source, "Testing") != "SO[Testing]" {
		t.Fatal("Formatting failed for: Source")
	}
}

func TestFormatting_SZ(t *testing.T) {
	if fmt.Sprintf(sgf.Size, 19) != "SZ[19]" {
		t.Fatal("Formatting failed for: Size")
	}
}

func TestFormatting_WP(t *testing.T) {
	if sgf.WhitePass != "W[]" {
		t.Fatal("Formatting failed for: WhitePass")
	}
}

func TestFormatting_WR(t *testing.T) {
	if fmt.Sprintf(sgf.WhiteRank, "9d") != "WR[9d]" {
		t.Fatal("Formatting failed for: WhiteRank")
	}
}

func TestHeader(t *testing.T) {
	header := sgf.Header()
	if header != "(;FF[4]GM[1]SZ[19]AP[GoGo]" {
		t.Fatalf("Failed SGF header generation. Got: %s", header)
	}
}

func TestEmitMove(t *testing.T) {
	move := sgf.EmitMove(1, 4, 4)
	if move != ";B[dd]" {
		t.Fatalf("Failed to EmitMove. Wanted: 'B[dd];', Got: '%s'", move)
	}
}
