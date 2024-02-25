package pgn

import (
	"chess-platform-analyzer/internal/chess"
	"strings"
)

type (
	PGN struct {
		Event          string   `json:"event"`
		Site           string   `json:"site"`
		Date           string   `json:"date"`
		White          string   `json:"white"`
		Black          string   `json:"black"`
		Result         string   `json:"result"`
		Variant        string   `json:"variant"`
		TimeControl    string   `json:"time_control"`
		ECO            string   `json:"eco"`
		GamePlainText  string   `json:"game_plain_text"`
		UCIFormatMoves []string `json:"game_algebraic_notation"`
	}
)

// parse a plain-text PGN to a slice of PGN struct
func ParseStringGames(games string) []PGN {
	var actualPGN PGN
	var splitedGames []PGN

	g := strings.Split(games, "\n")

	for _, v := range g {
		v = strings.Trim(v, "\t")
		if strings.HasPrefix(v, "[Event") && actualPGN.GamePlainText != "" {
			splitedGames = append(splitedGames, actualPGN)
			actualPGN = PGN{}
		}

		actualPGN.parsePlainTextPGNLine(v)
		actualPGN.parsePlainTextGameToUCIFormat()
	}

	splitedGames = append(splitedGames, actualPGN)
	return splitedGames
}

func (p *PGN) parsePlainTextPGNLine(line string) {
	l := strings.ReplaceAll(line, "\t", "")
	if !strings.HasPrefix(l, "[") {
		p.GamePlainText += l
		return
	}

	l = strings.ReplaceAll(line, "[", "")
	l = strings.ReplaceAll(l, "]", "")
	l = strings.ReplaceAll(l, `"`, "")
	l = strings.ReplaceAll(l, "\t", "")

	value, header := lookAnyHeader(l)
	if header == "" {
		return
	}

	insertHeaderInPGN(p, header, value)
}

func insertHeaderInPGN(p *PGN, header string, value string) {
	if header == "Event" {
		p.Event = value
		return
	}
	if header == "Site" {
		p.Site = value
		return
	}
	if header == "White" {
		p.White = value
		return
	}
	if header == "Black" {
		p.Black = value
		return
	}
	if header == "Date" {
		p.Date = value
		return
	}
	if header == "Result" {
		p.Result = value
		return
	}
	if header == "Variant" {
		p.Variant = value
		return
	}
	if header == "TimeControl" {
		p.TimeControl = value
		return
	}
	if header == "ECO" {
		p.ECO = value
		return
	}
}

// Search for all PGN headers that header parameter could be.
// If header is parseable to PGN struct, function returns header value and which field it is.
// If header is not parseable to PGN struct, function returns two empty strings.
func lookAnyHeader(header string) (string, string) {
	parseableHeaders := []string{"Event", "Site", "Date", "White", "Black", "Result", "Variant",
		"TimeControl", "ECO"}

	for _, v := range parseableHeaders {
		if strings.HasPrefix(header, v+" ") {
			return strings.Replace(header, v+" ", "", 1), v
		}
	}

	return "", ""
}

func (pgn *PGN) parsePlainTextGameToUCIFormat() {
	if pgn.GamePlainText == "" {
		return
	}

	moves := strings.Split(pgn.GamePlainText, " ")
	board := chess.NewBoard()
	for _, move := range moves {
		// if string contains a dot, string is not a move, it is move number (like "1.")
		if strings.Contains(move, ".") || move == " " || move == "\n" || move == "" {
			continue
		}

		if move == "1-0" || move == "0-1" || move == "1/2-1/2" {
			break
		}

		move = strings.ReplaceAll(move, "+", "")
		move = strings.ReplaceAll(move, "!", "")
		move = strings.ReplaceAll(move, "?", "")

		if move == "O-O" {
			if board.Turn == "w" {
				board.MakeMove("e1g1")
				continue
			}

			board.MakeMove("e8g8")
			continue
		}

		if move == "O-O-O" {
			if board.Turn == "w" {
				board.MakeMove("e1c1")
				continue
			}

			board.MakeMove("e8c8")
			continue
		}

		availableMoves := board.AvailableLegalMoves()

		// if move lenght is two, it is a pawn movement like e4
		if len(move) == 2 {
			board.MakeMove(getSameColumnMove(move, availableMoves))
			continue
		}

		// if move contains 'x' it is a capture
		if strings.Contains(move, "x") {
			board.MakeMove(getCaptureMove(move, availableMoves, board))
			continue
		}

		// if is not similar to another move type, it is a piece movement like Nf3
		board.MakeMove(getPieceMove(move, availableMoves, board))
	}

	pgn.UCIFormatMoves = board.MovesHistory
}

// returns the first move with equal column letter for pawn movements.
// e.g.: if move is e4 and available moves contains f3e4 and e2e4
// the functions will return e2e4.
func getSameColumnMove(move string, availableMoves []string) string {
	column := move[:1]
	for _, m := range availableMoves {
		if m[2:] != move {
			continue
		}

		if m[:1] == column {
			return m
		}
	}

	return ""
}

// returns a move from availableMoves after analize similarity with move parameter.
func getCaptureMove(move string, availableMoves []string, board chess.Board) string {
	// fxe4
	decomposedMove := strings.Split(move, "x")
	targetSquare := decomposedMove[1]

	var pieceOrColumn, origin string
	pieceOrColumn = decomposedMove[0]
	if len(decomposedMove[0]) == 2 {
		pieceOrColumn = string(decomposedMove[0][0])
		origin = string(decomposedMove[0][1])
	}

	for _, m := range availableMoves {
		if targetSquare != m[2:] {
			continue
		}

		// if move has a origin, it is a shared target move
		// like Nfxf3
		if origin != "" {
			originColumn := m[:1]
			originRow := m[1:2]
			if origin == originColumn || origin == originRow {
				return m
			}

			continue
		}

		poc := pieceOrColumn[0]
		// if first letter of movement is capitalized, it is a piece
		// like Nxf3
		if rune(poc) >= 'A' && rune(poc) <= 'Z' {
			piece := board.GetPieceAt(m[0:2])
			if !piece.IsColor(board.Turn) {
				continue
			}

			if strings.EqualFold(string(piece), string(poc)) {
				return m
			}
		}

		// if is not similar of any another moves type, it is a pawn capture
		// like gxf3
		if m[:1] == string(poc) {
			return m
		}
	}

	return ""
}

func getPieceMove(move string, availableMoves []string, board chess.Board) string {
	movePiece := move[:1]
	targetSquare := move[1:]
	origin := ""
	if len(move) == 4 {
		origin = move[1:2]
		targetSquare = move[2:]
	}
	for _, m := range availableMoves {
		if m[2:] != targetSquare {
			continue
		}

		if origin != "" && m[:1] != origin && m[1:2] != origin {
			continue
		}

		piece := board.GetPieceAt(m[:2])
		if strings.EqualFold(string(piece), movePiece) {
			return m
		}
	}

	return ""
}
