package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/JxBP/go-life/board"
	"github.com/JxBP/go-life/game"
)

func main() {
	width := flag.Int("width", 40, "width of the board")
	height := flag.Int("height", 40, "height of the board")
	interval := flag.String("interval", "1s", "timeout between generations")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	timeout, err := time.ParseDuration(*interval)
	if err != nil {
		log.Fatalf("err: bad interval: %v", err)
	}

	b := board.New(*width, *height)
	for x := range b.Buf {
		for y := range b.Buf[x] {
			if rand.Intn(7) < 2 {
				b.Buf[x][y] = true
			}
		}
	}

	g := game.WithBoard(b)

	for {
		fmt.Print("\033[2J\033[H")
		if *debug {
			fmt.Print(g.String())
		} else {
			fmt.Print(g.Display())
		}
		g.Step()
		time.Sleep(timeout)
	}
}
