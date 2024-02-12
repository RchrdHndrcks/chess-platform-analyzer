package main

import (
	"fmt"
	"time"

	"github.com/RchrdHndrcks/chess-analyzer/internal"
)

func main() {
	start := time.Now()

	user := "EddyRob"
	daysAgo := 5

	p := internal.LichessPlatform{}

	fmt.Println("Delayed time:", time.Now().Sub(start))
}
