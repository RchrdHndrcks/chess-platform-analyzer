package pgn

import "strings"

type (
	PGN struct {
		Event                 string `json:"event"`
		Site                  string `json:"site"`
		Date                  string `json:"date"`
		White                 string `json:"white"`
		Black                 string `json:"black"`
		Result                string `json:"result"`
		Variant               string `json:"variant"`
		TimeControl           string `json:"time_control"`
		ECO                   string `json:"eco"`
		GamePlainText         string `json:"game_plain_text"`
		GameAlgebraicNotation string `json:"game_algebraic_notation"`
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

func (pgn *PGN) parsePlainTextGameToAlgebraicNotation() {
	if pgn.GamePlainText == "" {
		return
	}

	//normalizedGames := strings.ReplaceAll(pgn.GamePlainText, " ", "")

}
