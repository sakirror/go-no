package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	tbl := [][]string{
		{"いやよ(怒)", "いやよ(照)"},
		{"no", "yes"},
	}

	rand.Seed(time.Now().UnixNano())
	for {
		i := 0
		if rand.Intn(10) == 8 {
			i = 1
		}
		fmt.Println(tbl[0][i])
		time.Sleep(10 * time.Millisecond)
	}
}
