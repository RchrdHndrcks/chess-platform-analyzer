package internal

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type (
	lichessPlatform struct{}
)

const url string = "https://lichess.org/api/games/user/%s?since=%d&perfType=blitz,rapid,classical"

func NewLichessPlatform() ChessPlatform {
	return lichessPlatform{}
}

// GetGamesFromManyDays
// user: user name from Lichess
// daysAgo: how many days ago games you want
// This function make a POST request to Lichess API looking for user games.
// After request the function return user games as string.
// Function return error if request or read response failed.
func (lichessPlatform) GetGamesFromManyDays(user string, daysAgo int) (string, error) {
	response, err := makeLichessRequest(user, daysAgo)
	if err != nil {
		return "", fmt.Errorf("error calling makeLichessRequest: %s", err)
	}

	_, games, err := readRequestResponse(response)
	if err != nil {
		return "", fmt.Errorf("error calling readRequestResponse: %s", err)
	}

	return string(games), nil
}

func makeLichessRequest(user string, daysAgo int) (*http.Response, error) {
	url := fmt.Sprintf(url, user, getTimeinTimestamp(daysAgo))
	return http.Get(url)
}

func getTimeinTimestamp(daysAgo int) int64 {
	return time.Now().AddDate(0, 0, -daysAgo).UnixMilli()
}

func readRequestResponse(resp *http.Response) (int, []byte, error) {
	bytesRead := 0
	buf := make([]byte, 1024*1024)
	game := []byte{}

	for {
		n, err := resp.Body.Read(buf)
		bytesRead += n

		game = append(game, buf...)

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, nil, err
		}
	}

	return bytesRead, game, nil
}
