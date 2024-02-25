package chess

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	Board struct {
		Board            [][]Piece
		Turn             string
		MovesCount       int
		HalfMoves        int
		InPasantSquare   string
		AvailableCastles string
		IsCheck          bool
		MovesHistory     []string
	}
)

var (
	columnLetter map[int]string = map[int]string{
		0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f", 6: "g", 7: "h",
	}
	rowNumber map[int]string = map[int]string{
		0: "8", 1: "7", 2: "6", 3: "5", 4: "4", 5: "3", 6: "2", 7: "1",
	}

	slidingPieces map[Piece]bool = map[Piece]bool{
		WBishop: true, WRook: true, BBishop: true, BRook: true,
		WQueen: true, BQueen: true, WKing: true, BKing: true,
	}

	whiteCoronationPieces []string = []string{"Q", "R", "N", "B"}
	blackCoronationPieces []string = []string{"q", "r", "n", "b"}
)

func NewBoard() Board {
	b := Board{}
	b.TranslateFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	return b
}

func (board *Board) MakeMove(movement string) {
	originSquare, targetSquare, coronationPiece := extractSquaresFromMovement(movement)

	xOrigin, yOrigin := generateXYFromSquare(originSquare)
	xTarget, yTarget := generateXYFromSquare(targetSquare)

	board.MovesHistory = append(board.MovesHistory, movement)
	board.move(xOrigin, yOrigin, xTarget, yTarget, Piece(coronationPiece))
}

// calculate all available legal moves for board.Turn color.
// return: a slice with content if legal moves exist, empty slice if is stalemate or nil if is checkmate
func (board Board) AvailableLegalMoves() []string {
	king := WKing
	if board.Turn == "b" {
		king = BKing
	}

	availableMoves := board.availableMoves()
	legalMovements := []string{}

	for _, move := range availableMoves {
		legalMove := true

		boardPlayground := board.copyBoard()
		boardPlayground.MakeMove(move)
		opponentAvailableMoves := boardPlayground.availableMoves()
		for _, opMove := range opponentAvailableMoves {
			_, opTargetSquare, _ := extractSquaresFromMovement(opMove)

			kingSquare := boardPlayground.WhereIs(king)[0]
			// if opponent has a movement where target square is my king
			// my move was with my king under attack, ergo ilegal.
			if opTargetSquare == kingSquare {
				legalMove = false
				break
			}

			originSquare, targetSquare, _ := extractSquaresFromMovement(move)
			x, y := generateXYFromSquare(originSquare)
			if isCastleMovement(board.Board[y][x], originSquare, targetSquare) &&
				(board.isCheck() || targetSquareIsInKingWayInCastle(targetSquare)) {
				legalMove = false
				break
			}

			boardPlayground = board.copyBoard()
			boardPlayground.MakeMove(move)
		}

		if legalMove {
			legalMovements = append(legalMovements, move)
		}
	}

	// if are not legal movements king could be in mate or stalemate
	if len(legalMovements) == 0 {
		if board.isCheck() {
			return nil
		}
	}

	return legalMovements
}

// returns FEN from position
func (b Board) FEN() string {
	fen := ""
	emptySquares := 0
	for i := range b.Board {
		f := ""
		for _, square := range b.Board[i] {
			if square == "" {
				emptySquares++
				continue
			}

			appending := fmt.Sprintf("%d%s", emptySquares, square)
			if emptySquares == 0 {
				appending = string(square)
			}
			emptySquares = 0

			f += appending
		}

		if emptySquares != 0 {
			f += strconv.Itoa(emptySquares)
			emptySquares = 0
		}

		fen += f + "/"
	}

	return fmt.Sprintf("%s %s %s %s %d %d", strings.Trim(fen, "/"), b.Turn, b.AvailableCastles,
		b.InPasantSquare, b.HalfMoves, b.MovesCount)
}

// given a FEN function translate it to a board.
// this function will change your board.
func (board *Board) TranslateFEN(FEN string) error {
	rows := strings.Split(FEN, "/")
	if len(rows) != 8 {
		return fmt.Errorf("invalid FEN")
	}

	b := [][]Piece{}
	for x := 0; x < 8; x++ {
		row := []Piece{}
		var additionalInformation []string
		if x == 7 {
			additionalInformation = strings.Split(rows[x], " ")
			rows[x] = additionalInformation[0]
			additionalInformation = additionalInformation[1:]
		}

		for y := 0; y < 8; y++ {
			c := string(rows[x][0])
			rows[x] = rows[x][1:]

			n, err := strconv.Atoi(c)
			if err != nil {
				row = append(row, Piece(c))
				continue
			}

			row = append(row, "")
			n--
			if n > 0 {
				rows[x] = fmt.Sprintf("%d%s", n, rows[x])
			}
		}

		b = append(b, row)

		if additionalInformation != nil {
			err := board.setAdditionalProperties(additionalInformation)
			if err != nil {
				return err
			}
		}
	}

	board.Board = b
	return nil
}

func (b Board) WhereIs(p Piece) []string {
	squares := []string{}
	for y, row := range b.Board {
		for x, piece := range row {
			if piece != p {
				continue
			}

			squares = append(squares, generateSquare(x, y))
		}
	}

	return squares
}

// give all possible moves (legal or not)
func (b Board) availableMoves() []string {
	moves := ""
	for y, row := range b.Board {
		for x, piece := range row {
			if piece == "" {
				continue
			}

			if !piece.IsColor(b.Turn) {
				continue
			}

			specialMoves := b.generateSpecialMoves(x, y)
			if specialMoves != "" {
				moves = joinStrings(moves, specialMoves)
			}

			isSlidingPiece := slidingPieces[piece]
			if isSlidingPiece {
				slidingMoves := b.generateSlidingPiecesMoves(x, y)
				moves = joinStrings(moves, slidingMoves)
				continue
			}

			if piece == WKnight || piece == BKnight {
				knightMoves := b.generateKnightMoves(x, y)
				moves = joinStrings(moves, knightMoves)
				continue
			}

			if piece == WPawn || piece == BPawn {
				pawnMoves := b.generatePawnMoves(x, y)
				moves = joinStrings(moves, pawnMoves)
			}
		}
	}

	m := strings.Trim(moves, " ")
	if m == "" {
		return nil
	}

	return strings.Split(m, " ")
}

// generate different pieces special moves, except pawn transformation
// because this move is not different to normal pawn movement
func (board Board) generateSpecialMoves(x, y int) string {
	piece := board.Board[y][x]

	if (piece == WKing || piece == BKing) && board.AvailableCastles != "" {
		return strings.Trim(board.possibleCastles(x, y), " ")
	}

	if piece == WPawn || piece == BPawn {
		return strings.Trim(board.possibleInPassant(x, y), " ")
	}

	return ""
}

func (board Board) generatePawnMoves(x, y int) string {
	pawn := board.Board[y][x]

	isWhiteColor := pawn.IsColor("w")

	yOffset := -1
	if !isWhiteColor {
		yOffset = 1
	}

	squareCount := 1
	if isWhiteColor && y == 6 || !isWhiteColor && y == 1 {
		squareCount++
	}

	moves := ""
	for i := 1; i <= squareCount; i++ {
		yTarget := y + yOffset*i
		if !board.isInsideLimit(x, yTarget) {
			break
		}

		target := board.Board[yTarget][x]
		if target != "" {
			break
		}

		// if target square is board limit, the movement is a pawn coronation
		if yTarget == 0 || yTarget == 7 {
			possibleCoronation := whiteCoronationPieces
			if !isWhiteColor {
				possibleCoronation = blackCoronationPieces
			}

			for _, p := range possibleCoronation {
				moves += generateTargetMovement(x, y, x, yTarget) + p
			}

			continue
		}

		moves += generateTargetMovement(x, y, x, yTarget)
	}

	// after look in front of pawn, look for possible captures
	xTarget := x + 1
	yTarget := y + yOffset
	if board.isInsideLimit(xTarget, yTarget) {
		targetCapture := board.Board[yTarget][xTarget]
		if targetCapture != "" && !targetCapture.AreSameColor(pawn) {
			moves += generateTargetMovement(x, y, xTarget, yTarget)
		}
	}

	xTarget = x - 1
	if board.isInsideLimit(xTarget, yTarget) {
		targetCapture := board.Board[yTarget][xTarget]
		if targetCapture != "" && !targetCapture.AreSameColor(pawn) {
			moves += generateTargetMovement(x, y, xTarget, yTarget)
		}
	}

	return strings.Trim(moves, " ")
}

func (board Board) generateKnightMoves(x, y int) string {
	knight := board.Board[y][x]
	offsets := map[int][]int{
		1:  {2, -2},
		-1: {2, -2},
		2:  {1, -1},
		-2: {1, -1},
	}

	moves := ""
	for xOffset, offsetsY := range offsets {
		for _, yOffset := range offsetsY {
			xTarget := x + xOffset
			yTarget := y + yOffset
			if !board.isInsideLimit(xTarget, yTarget) {
				continue
			}

			if board.Board[yTarget][xTarget].AreSameColor(knight) {
				continue
			}

			moves += generateTargetMovement(x, y, xTarget, yTarget)
		}
	}

	return strings.Trim(moves, " ")
}

func (board Board) generateSlidingPiecesMoves(x, y int) string {
	moves := ""
	piece := board.Board[y][x]

	directionsCount := generateDirectionsCount(piece)
	for d := 0; d < directionsCount; d++ {
		offsetFunc := generateOffsetFunction(piece)
		if offsetFunc == nil {
			continue
		}

		xOffset, yOffset := offsetFunc(d)
		moves += board.lookForMovesAtDirection(x, xOffset, y, yOffset)
	}

	return strings.Trim(moves, " ")
}

func (board Board) lookForMovesAtDirection(x, xOffset, y, yOffset int) string {
	moves := ""
	piece := board.Board[y][x]
	xTarget := x + xOffset
	yTarget := y + yOffset
	for board.isInsideLimit(xTarget, yTarget) {
		targetPiece := board.Board[yTarget][xTarget]

		// direction blocked by friendly piece
		if piece.AreSameColor(targetPiece) {
			break
		}

		// if target square is blocked by an opponent piece or is empty
		// it is a valid square to move
		moves += generateTargetMovement(x, y, xTarget, yTarget)

		// direction blocked by opponent piece
		if targetPiece != "" {
			break
		}

		if piece == WKing || piece == BKing {
			break
		}

		xTarget += xOffset
		yTarget += yOffset
	}

	return moves
}

func (board Board) possibleCastles(x, y int) string {
	king := board.Board[y][x]
	moves := ""
	for _, v := range board.AvailableCastles {
		side := string(v)
		possibleCastleSide := Piece(side)
		if !king.AreSameColor(possibleCastleSide) {
			continue
		}

		side = strings.ToLower(side)
		offset := 1
		if side == "q" {
			offset = -1
		}

		if board.Board[y][x+offset] != "" || board.Board[y][x+offset*2] != "" {
			continue
		}

		moves += generateTargetMovement(x, y, x+offset*2, y)
	}

	return moves
}

func (board Board) possibleInPassant(x, y int) string {
	pawn := board.Board[y][x]
	xOffsets := []int{-1, 1}
	yOffset := -1
	if pawn.IsColor("b") {
		yOffset = 1
	}

	for _, xOffset := range xOffsets {
		xTarget := x + xOffset
		yTarget := y + yOffset
		targetSquare := fmt.Sprintf("%s%s", columnLetter[xTarget], rowNumber[yTarget])
		if board.InPasantSquare == targetSquare {
			return generateTargetMovement(x, y, xTarget, yTarget)
		}
	}

	return ""
}

func (board *Board) setAdditionalProperties(additionalInformation []string) error {
	board.Turn = additionalInformation[0]
	board.AvailableCastles = additionalInformation[1]
	board.InPasantSquare = additionalInformation[2]

	halfMoves, err := strconv.Atoi(additionalInformation[3])
	if err != nil {
		return fmt.Errorf("invalid FEN")
	}

	movesCount, err := strconv.Atoi(additionalInformation[4])
	if err != nil {
		return fmt.Errorf("invalid FEN")
	}

	board.HalfMoves = halfMoves
	board.MovesCount = movesCount
	return nil
}

func (board *Board) haveLostCastle(x, y, xTarget, yTarget int) {
	if board.AvailableCastles == "" {
		return
	}

	p := board.Board[y][x]
	if p == WKing || p == BKing || p == WRook || p == BRook {
		possibleCastles := ""
		for _, c := range board.AvailableCastles {
			str := string(c)
			if !p.AreSameColor(Piece(str)) {
				possibleCastles += str
			}
		}
	}

	// a8 square
	if xTarget == 0 && yTarget == 0 {
		board.AvailableCastles = strings.Replace(board.AvailableCastles, "q", "", 1)
		return
	}

	// h8 square
	if xTarget == 7 && yTarget == 0 {
		board.AvailableCastles = strings.Replace(board.AvailableCastles, "k", "", 1)
		return
	}

	// h1 square
	if xTarget == 7 && yTarget == 7 {
		board.AvailableCastles = strings.Replace(board.AvailableCastles, "K", "", 1)
		return
	}

	// a1 square
	if xTarget == 0 && yTarget == 7 {
		board.AvailableCastles = strings.Replace(board.AvailableCastles, "Q", "", 1)
	}
}

func (board *Board) move(x, y, xTarget, yTarget int, coronationPiece Piece) {
	p := board.Board[y][x]

	board.Turn = "b"
	if p.IsColor("b") {
		board.MovesCount++
		board.Turn = "w"
	}

	board.HalfMoves++
	if p == WPawn || p == BPawn {
		board.HalfMoves = 0
		if y == 6 && yTarget == 4 || y == 1 && yTarget == 3 {
			offset := 1
			if p.IsColor("b") {
				offset = -1
			}

			board.InPasantSquare = generateSquare(xTarget, yTarget+offset)
		}
	}

	board.haveLostCastle(x, y, xTarget, yTarget)

	if coronationPiece != "" {
		p = Piece(coronationPiece)
	}

	isCastleMove := isCastleMovement(p, generateSquare(x, y), generateSquare(xTarget, yTarget))

	board.Board[y][x] = ""
	board.Board[yTarget][xTarget] = p

	if isCastleMove {
		rX, rY, rXTarget, rYTarget := getCastleRookPosition(xTarget, yTarget)
		board.Board[rYTarget][rXTarget] = board.Board[rY][rX]
		board.Board[rY][rX] = ""
	}
}

func (board Board) isCheck() bool {
	king := WKing
	b := board.copyBoard()
	b.Turn = "b"
	if board.Turn == "b" {
		b.Turn = "w"
		king = BKing
	}

	b.InPasantSquare = ""

	kingSquare := board.WhereIs(king)[0]
	opponentMoves := b.AvailableLegalMoves()
	for _, move := range opponentMoves {
		_, targetSquare, _ := extractSquaresFromMovement(move)

		if targetSquare == kingSquare {
			return true
		}
	}

	return false
}

func (Board) parseToAlgebraicNotation(game string) string {

	return ""
}

func (Board) isInsideLimit(x, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func (board Board) copyBoard() Board {
	c := Board{}
	c.TranslateFEN(board.FEN())
	return c
}
