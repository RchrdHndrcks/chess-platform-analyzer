package main

import (
	"fmt"
	"time"

	"chess-platform-analyzer/internal"
)

func main() {
	start := time.Now()

	user := "EddyRob"
	daysAgo := 5

	p := internal.LichessPlatform{}
	games, err := p.GetGamesFromManyDays(user, daysAgo)
	if err != nil {
		fmt.Printf("error calling GetGamesFromManyDays: %s", err)
		return
	}

	fmt.Println(games)
	fmt.Println("Delayed time:", time.Now().Sub(start))
}
