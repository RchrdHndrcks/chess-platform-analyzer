package pgn

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEachLineOfString(t *testing.T) {
	assert := assert.New(t)
	game := `[Event "Rated Blitz game"]
	[Site "https://lichess.org/R2Mc2Oi3"]
	[Date "2024.02.09"]
	[White "EddyRob"]
	[Black "Steevie"]
	[Result "1-0"]
	[UTCDate "2024.02.09"]
	[UTCTime "19:51:46"]
	[WhiteElo "2048"]
	[BlackElo "2030"]
	[WhiteRatingDiff "+5"]
	[BlackRatingDiff "-6"]
	[Variant "Standard"]
	[TimeControl "180+2"]
	[ECO "B28"]
	[Termination "Normal"]
	
	1. e4 c5 2. Nf3 a6 3. d4 cxd4 4. Nxd4 e5 5. Nb3 Nf6 6. Nc3 Bb4 7. f3 d5 8. exd5 Nxd5 9. Bd2 Nxc3 10. Bxc3 Bxc3+ 11. bxc3 Qxd1+ 12. Rxd1 Nc6 13. Nc5 b6 14. Ne4 Be6 15. Nd6+ Ke7 16. Bc4 Rhd8 17. Bxe6 Kxe6 18. Nc4 Rxd1+ 19. Kxd1 b5 20. Ne3 f5 21. Ke2 Rd8 22. Rb1 f4 23. Nd1 Rd5 24. Nf2 Na5 25. a4 bxa4 26. Rb6+ Rd6 27. Rb4 Nc6 28. Rxa4 a5 29. Ne4 Rd5 30. c4 Rd7 31. Nc5+ Kd6 32. Nxd7 Kxd7 33. Kd3 Kd6 34. Ke4 Kc5 35. Kf5 Kd4 36. Ke6 Ke3 37. Kd6 Nd4 38. Kxe5 Nc6+ 39. Kd6 
	Nd4 40. c5 Kf2 41. Rxd4 1-0`
	pgn := PGN{}
	lines := strings.Split(game, "\n")

	for _, v := range lines {
		pgn.parsePlainTextPGNLine(v)
	}

	assert.Equal("Rated Blitz game", pgn.Event)
	assert.Equal("https://lichess.org/R2Mc2Oi3", pgn.Site)
	assert.Equal("2024.02.09", pgn.Date)
	assert.Equal("EddyRob", pgn.White)
	assert.Equal("Steevie", pgn.Black)
	assert.Equal("1-0", pgn.Result)
	assert.Equal("Standard", pgn.Variant)
	assert.Equal("180+2", pgn.TimeControl)
	assert.Equal("B28", pgn.ECO)
	assert.Equal("1. e4 c5 2. Nf3 a6 3. d4 cxd4 4. Nxd4 e5 5. Nb3 Nf6 6. Nc3 Bb4 7. f3 d5 8. exd5 Nxd5 9. Bd2 Nxc3 10. Bxc3 Bxc3+ 11. bxc3 Qxd1+ 12. Rxd1 Nc6 13. Nc5 b6 14. Ne4 Be6 15. Nd6+ Ke7 16. Bc4 Rhd8 17. Bxe6 Kxe6 18. Nc4 Rxd1+ 19. Kxd1 b5 20. Ne3 f5 21. Ke2 Rd8 22. Rb1 f4 23. Nd1 Rd5 24. Nf2 Na5 25. a4 bxa4 26. Rb6+ Rd6 27. Rb4 Nc6 28. Rxa4 a5 29. Ne4 Rd5 30. c4 Rd7 31. Nc5+ Kd6 32. Nxd7 Kxd7 33. Kd3 Kd6 34. Ke4 Kc5 35. Kf5 Kd4 36. Ke6 Ke3 37. Kd6 Nd4 38. Kxe5 Nc6+ 39. Kd6 Nd4 40. c5 Kf2 41. Rxd4 1-0", pgn.GamePlainText)
}

func TestParseStringGames(t *testing.T) {
	assert := assert.New(t)
	games := `[Event "Rated Blitz game"]
	[Site "https://lichess.org/R2Mc2Oi3"]
	[Date "2024.02.09"]
	[White "EddyRob"]
	[Black "Steevie"]
	[Result "1-0"]
	[UTCDate "2024.02.09"]
	[UTCTime "19:51:46"]
	[WhiteElo "2048"]
	[BlackElo "2030"]
	[WhiteRatingDiff "+5"]
	[BlackRatingDiff "-6"]
	[Variant "Standard"]
	[TimeControl "180+2"]
	[ECO "B28"]
	[Termination "Normal"]
	
	1. e4 c5 2. Nf3 a6 3. d4 cxd4 4. Nxd4 e5 5. Nb3 Nf6 6. Nc3 Bb4 7. f3 d5 8. exd5 Nxd5 9. Bd2 Nxc3 10. Bxc3 Bxc3+ 11. bxc3 Qxd1+ 12. Rxd1 Nc6 13. Nc5 b6 14. Ne4 Be6 15. Nd6+ Ke7 16. Bc4 Rhd8 17. Bxe6 Kxe6 18. Nc4 Rxd1+ 19. Kxd1 b5 20. Ne3 f5 21. Ke2 Rd8 22. Rb1 f4 23. Nd1 Rd5 24. Nf2 Na5 25. a4 bxa4 26. Rb6+ Rd6 27. Rb4 Nc6 28. Rxa4 a5 29. Ne4 Rd5 30. c4 Rd7 31. Nc5+ Kd6 32. Nxd7 Kxd7 33. Kd3 Kd6 34. Ke4 Kc5 35. Kf5 Kd4 36. Ke6 Ke3 37. Kd6 Nd4 38. Kxe5 Nc6+ 39. Kd6 Nd4 40. c5 Kf2 41. Rxd4 1-0
	
	
	[Event "Rated Blitz game"]
	[Site "https://lichess.org/4wybg79d"]
	[Date "2024.02.09"]
	[White "kakaobohne"]
	[Black "EddyRob"]
	[Result "0-1"]
	[UTCDate "2024.02.09"]
	[UTCTime "18:00:23"]
	[WhiteElo "2020"]
	[BlackElo "2043"]
	[WhiteRatingDiff "-5"]
	[BlackRatingDiff "+5"]
	[Variant "Standard"]
	[TimeControl "180+2"]
	[ECO "E76"]
	[Termination "Normal"]
	
	1. d4 Nf6 2. c4 c5 3. d5 g6 4. Nc3 d6 5. e4 Bg7 6. f4 O-O 7. Nf3 a6 8. e5 Nfd7 9. e6 Nf6 10. exf7+ Rxf7 11. Ng5 Bg4 12. Nxf7 Kxf7 13. Qb3 Qd7 14. h3 Bf5 15. Be2 e6 16. g4 Be4 17. Nxe4 Nxe4 18. f5 exf5 19. gxf5 gxf5 20. Bh5+ Kg8 21. O-O Qe7 22. Bf3 Bd4+ 23. Kh2 Qe5+ 24. Kh1 Nf2+ 25. Kg2 Qg7+ 26. Kh2 Be5+ 0-1`

	pgns := ParseStringGames(games)

	assert.Len(pgns, 2)
	assert.Equal("Rated Blitz game", pgns[0].Event)
	assert.Equal("https://lichess.org/R2Mc2Oi3", pgns[0].Site)
	assert.Equal("2024.02.09", pgns[0].Date)
	assert.Equal("EddyRob", pgns[0].White)
	assert.Equal("Steevie", pgns[0].Black)
	assert.Equal("1-0", pgns[0].Result)
	assert.Equal("Standard", pgns[0].Variant)
	assert.Equal("180+2", pgns[0].TimeControl)
	assert.Equal("B28", pgns[0].ECO)
	assert.Equal("1. e4 c5 2. Nf3 a6 3. d4 cxd4 4. Nxd4 e5 5. Nb3 Nf6 6. Nc3 Bb4 7. f3 d5 8. exd5 Nxd5 9. Bd2 Nxc3 10. Bxc3 Bxc3+ 11. bxc3 Qxd1+ 12. Rxd1 Nc6 13. Nc5 b6 14. Ne4 Be6 15. Nd6+ Ke7 16. Bc4 Rhd8 17. Bxe6 Kxe6 18. Nc4 Rxd1+ 19. Kxd1 b5 20. Ne3 f5 21. Ke2 Rd8 22. Rb1 f4 23. Nd1 Rd5 24. Nf2 Na5 25. a4 bxa4 26. Rb6+ Rd6 27. Rb4 Nc6 28. Rxa4 a5 29. Ne4 Rd5 30. c4 Rd7 31. Nc5+ Kd6 32. Nxd7 Kxd7 33. Kd3 Kd6 34. Ke4 Kc5 35. Kf5 Kd4 36. Ke6 Ke3 37. Kd6 Nd4 38. Kxe5 Nc6+ 39. Kd6 Nd4 40. c5 Kf2 41. Rxd4 1-0", pgns[0].GamePlainText)
	assert.Equal([]string{"e2e4", "c7c5", "g1f3", "a7a6", "d2d4", "c5d4", "f3d4", "e7e5", "d4b3", "g8f6", "b1c3", "f8b4",
		"f2f3", "d7d5", "e4d5", "f6d5", "c1d2", "d5c3", "d2c3", "b4c3", "b2c3", "d8d1", "a1d1", "b8c6", "b3c5",
		"b7b6", "c5e4", "c8e6", "e4d6", "e8e7", "f1c4", "h8d8", "c4e6", "e7e6", "d6c4", "d8d1", "e1d1", "b6b5",
		"c4e3", "f7f5", "d1e2", "a8d8", "h1b1", "f5f4", "e3d1", "d8d5", "d1f2", "c6a5", "a2a4", "b5a4", "b1b6",
		"d5d6", "b6b4", "a5c6", "b4a4", "a6a5", "f2e4", "d6d5", "c3c4", "d5d7", "e4c5", "e6d6", "c5d7", "d6d7",
		"e2d3", "d7d6", "d3e4", "d6c5", "e4f5", "c5d4", "f5e6", "d4e3", "e6d6", "c6d4", "d6e5", "d4c6", "e5d6",
		"c6d4", "c4c5", "e3f2", "a4d4"}, pgns[0].UCIFormatMoves)

	assert.Equal("Rated Blitz game", pgns[1].Event)
	assert.Equal("https://lichess.org/4wybg79d", pgns[1].Site)
	assert.Equal("2024.02.09", pgns[1].Date)
	assert.Equal("kakaobohne", pgns[1].White)
	assert.Equal("EddyRob", pgns[1].Black)
	assert.Equal("0-1", pgns[1].Result)
	assert.Equal("Standard", pgns[1].Variant)
	assert.Equal("180+2", pgns[1].TimeControl)
	assert.Equal("E76", pgns[1].ECO)
	assert.Equal("1. d4 Nf6 2. c4 c5 3. d5 g6 4. Nc3 d6 5. e4 Bg7 6. f4 O-O 7. Nf3 a6 8. e5 Nfd7 9. e6 Nf6 10. exf7+ Rxf7 11. Ng5 Bg4 12. Nxf7 Kxf7 13. Qb3 Qd7 14. h3 Bf5 15. Be2 e6 16. g4 Be4 17. Nxe4 Nxe4 18. f5 exf5 19. gxf5 gxf5 20. Bh5+ Kg8 21. O-O Qe7 22. Bf3 Bd4+ 23. Kh2 Qe5+ 24. Kh1 Nf2+ 25. Kg2 Qg7+ 26. Kh2 Be5+ 0-1", pgns[1].GamePlainText)
	assert.Equal([]string{"d2d4", "g8f6", "c2c4", "c7c5", "d4d5", "g7g6", "b1c3", "d7d6", "e2e4", "f8g7", "f2f4",
		"e8g8", "g1f3", "a7a6", "e4e5", "f6d7", "e5e6", "d7f6", "e6f7", "f8f7", "f3g5", "c8g4", "g5f7", "g8f7",
		"d1b3", "d8d7", "h2h3", "g4f5", "f1e2", "e7e6", "g2g4", "f5e4", "c3e4", "f6e4", "f4f5", "e6f5", "g4f5",
		"g6f5", "e2h5", "f7g8", "e1g1", "d7e7", "h5f3", "g7d4", "g1h2", "e7e5", "h2h1", "e4f2", "h1g2", "e5g7",
		"g2h2", "d4e5"}, pgns[1].UCIFormatMoves)
}
