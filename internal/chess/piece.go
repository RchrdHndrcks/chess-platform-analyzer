package chess

type Piece string

const (
	WPawn   Piece = "P"
	WRook   Piece = "R"
	WKing   Piece = "K"
	WQueen  Piece = "Q"
	WBishop Piece = "B"
	WKnight Piece = "N"
	BPawn   Piece = "p"
	BRook   Piece = "r"
	BKing   Piece = "k"
	BQueen  Piece = "q"
	BBishop Piece = "b"
	BKnight Piece = "n"
)

func (p Piece) AreSameColor(P Piece) bool {
	p1 := string(p)
	p2 := string(P)
	if len(p1) == 0 || len(p2) == 0 {
		return false
	}

	rp := rune(p1[0])
	rP := rune(p2[0])

	return (rp > 'A' && rp < 'Z' && rP > 'A' && rP < 'Z') ||
		(rp > 'a' && rp < 'z' && rP > 'a' && rP < 'z')
}

// returns true if parameter color is same that piece
// color must be 'w' or 'b' like in FEN
func (p Piece) IsColor(color string) bool {
	if p == "" {
		return false
	}

	c := rune(string(p)[0])
	return (c > 'A' && c < 'Z' && color == "w") || (c > 'a' && c < 'z' && color == "b")
}
