package chess

import (
	"fmt"
	"math"
	"strings"
)

func generateDirectionsCount(piece Piece) int {
	p := strings.ToLower(string(piece))
	switch p {
	case "b", "r":
		return 4
	case "k", "q":
		return 8
	}

	return 0
}

func generateOffsetFunction(piece Piece) func(int) (int, int) {
	p := strings.ToLower(string(piece))
	switch p {
	case "b":
		return func(direction int) (int, int) {
			xOffset := 1
			if direction > 1 {
				xOffset = -1
			}
			yOffset := int(math.Pow(-1, float64(direction)))

			return xOffset, yOffset
		}
	case "r":
		return func(direction int) (int, int) {
			xOffset := direction % 2
			if direction == 3 {
				xOffset *= -1
			}

			yOffset := (direction + 1) % 2
			if direction == 2 {
				yOffset *= -1
			}

			return xOffset, yOffset
		}
	case "q", "k":
		return func(direction int) (int, int) {
			xOffset := -1
			if direction%4 == 0 {
				xOffset = 0
			} else if direction < 4 {
				xOffset = 1
			}

			yOffset := 1
			if direction == 0 || direction == 1 || direction == 7 {
				yOffset = -1
			} else if direction%2 == 0 && direction != 4 {
				yOffset = 0
			}

			return xOffset, yOffset
		}
	}

	return nil
}

func generateTargetMovement(x, y, xTarget, yTarget int) string {
	return fmt.Sprintf(" %s%s%s%s", columnLetter[x], rowNumber[y], columnLetter[xTarget], rowNumber[yTarget])
}

func generateSquare(x, y int) string {
	return fmt.Sprintf("%s%s", columnLetter[x], rowNumber[y])
}

func joinStrings(str1, str2 string) string {
	if str2 == "" {
		return str1
	}

	return str1 + " " + str2
}

func generateXYFromSquare(square string) (int, int) {
	var x, y int

	col := square[:1]
	for xCol, v := range columnLetter {
		if col == v {
			x = xCol
			break
		}
	}

	row := square[1:]
	for yRow, v := range rowNumber {
		if v == row {
			y = yRow
			break
		}
	}

	return x, y
}

func isCastleMovement(piece Piece, originSquare, targetSquare string) bool {
	if piece != WKing && piece != BKing {
		return false
	}

	if originSquare != "e1" && originSquare != "e8" {
		return false
	}

	if targetSquare != "g1" && targetSquare != "c1" && targetSquare != "g8" && targetSquare != "c8" {
		return false
	}

	return true
}

func targetSquareIsInKingWayInCastle(targetSquare string) bool {
	return targetSquare == "f1" || targetSquare == "d1" || targetSquare == "f8" || targetSquare == "d8"
}

func extractSquaresFromMovement(movement string) (string, string, string) {
	coronationPiece := ""
	originSquare := movement[:2]
	targetSquare := movement[2:]
	if len(targetSquare) == 3 {
		coronationPiece = targetSquare[2:]
		targetSquare = targetSquare[:2]
	}
	return originSquare, targetSquare, coronationPiece
}

func copyBoard(b Board) Board {
	c := Board{}
	c.TranslateFEN(b.FEN())
	return c
}
