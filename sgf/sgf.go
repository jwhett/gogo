// sgf describes the SGF file formatting and provides functions
// for working with SGF data. SGF FF[4] is the current file format
// as of this writing. See https://www.red-bean.com/sgf/ and
// https://www.red-bean.com/sgf/proplist_t.html
package sgf

const (
    // Application describes the application used
    // to generate the SGF file.
    Application string = "AP[%s]"
    // BlackPass shows that the Black player passed.
    BlackPass = "B[]"
    // BlackRank describes the player's rank.
    BlackRank = "BR[%s]"
    // Comment is a free-form comment string.
    Comment = "C[%s]"
    // DateTime is the date and time of the game.
    DateTime = "DT[%s]"
    // FileFormat describes the SGF version.
    FileFormat = "FF[%d]"
    // Game is the type of game being represented. GM[1] is Go/Baduk.
    Game = "GM[1]"
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
    // WhitePass shows that the White player passed.
    WhitePass = "W[]"
    // WhiteRank is the White player's rank.
    WhiteRank = "WR[%s]"
)
