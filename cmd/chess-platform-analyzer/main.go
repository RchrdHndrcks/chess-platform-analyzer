package main

import (
	"fmt"
	"time"

	platforms "chess-platform-analyzer/internal/platforms"
)

func main() {
	start := time.Now()

	user := "EddyRob"
	daysAgo := 5

	p := platforms.NewLichessPlatform()
	_, err := p.GetGamesFromManyDays(user, daysAgo)
	if err != nil {
		fmt.Printf("error calling GetGamesFromManyDays: %s", err)
		return
	}

	fmt.Println("Delayed time:", time.Now().Sub(start))
}
