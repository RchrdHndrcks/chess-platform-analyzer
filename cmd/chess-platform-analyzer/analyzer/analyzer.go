package analyzers

import "chess-platform-analyzer/internal/pgn"

type (
	GameStats struct {
		WinPercentage   float64
		LostPercentage  float64
		ECOPercentage   map[string]float64
		MostPlayedColor string
	}

	ECORegister map[string]float64
)

func ObtainGameStats(user string, games []pgn.PGN) GameStats {
	if len(games) == 0 {
		return GameStats{}
	}

	var wins float64
	var ECOPercentage ECORegister
	var gamesPlayedWithWhiteColor int
	for _, v := range games {
		wins += parseResult(user, v)
		ECOPercentage.registerGame(v)
		gamesPlayedWithWhiteColor += gamePlayedWithWhite(user, v)
	}

	winPercentage := wins / float64(len(games))
	gamesPlayedWithBlackColor := len(games) - gamesPlayedWithWhiteColor
	mostUsedColor := "WHITE"
	if gamesPlayedWithWhiteColor < gamesPlayedWithBlackColor {
		mostUsedColor = "BLACK"
	}

	return GameStats{
		WinPercentage:   winPercentage,
		LostPercentage:  100 - winPercentage,
		ECOPercentage:   ECOPercentage,
		MostPlayedColor: mostUsedColor,
	}
}

func parseResult(user string, p pgn.PGN) float64 {
	if p.Result == "1/2-1/2" {
		return 0.5
	}

	if p.White == user && p.Result == "1-0" {
		return 1
	}

	if p.Black == user && p.Result == "0-1" {
		return 1
	}

	return 0
}

func (r *ECORegister) registerGame(p pgn.PGN) {
	(*r)[p.ECO] += 1
}

func gamePlayedWithWhite(user string, p pgn.PGN) int {
	if p.White == user {
		return 1
	}

	return 0
}
