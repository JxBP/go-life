package main

import (
	"fmt"
	"time"

	"github.com/JxBP/go-life/board"
	"github.com/JxBP/go-life/game"
)

func main() {
	b := board.New(50, 20)
	b.Buf[1][3] = true
	b.Buf[2][1] = true
	b.Buf[2][3] = true
	b.Buf[3][2] = true
	b.Buf[3][3] = true

	g := game.WithBoard(b)

	for {
		fmt.Print("\033[2J\033[H")
		fmt.Print(g.String())
		g.Step()
		time.Sleep(time.Second * 1 / 10)
	}
}
