package internal

type (
	ChessPlatform interface {
		GetGamesFromManyDays(user string, daysAgo int) (string, error)
	}
)
