package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	tbl := [][]string{
		{"いやよ(怒)", "いやよ(照)"},
		{"no", "yes"},
	}

	if len(os.Args) >= 2 {
		tbl[0][0] = os.Args[1]
	}

	rand.Seed(time.Now().UnixNano())
	for {
		i := 0
		if rand.Intn(10) == 8 {
			i = 1
		}
		fmt.Println(tbl[0][i])
		time.Sleep(100 * time.Millisecond)
	}
}
