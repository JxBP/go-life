package game

import (
	"fmt"
	"strings"

	"github.com/JxBP/go-life/board"
)

type Game struct {
	curr, next *board.Board
}

func New(width int, height int) *Game {
	return &Game{
		board.New(width, height),
		board.New(width, height),
	}
}

func WithBoard(b *board.Board) *Game {
	return &Game{
		b,
		board.New(b.Width(), b.Height()),
	}
}

// TODO: Add a test for this
func NextState(neighbours int, isAlive bool) bool {
	return neighbours == 3 || (neighbours == 2 && isAlive)
}

func (g *Game) Step() {
	for x := range g.curr.Buf {
		for y := range g.curr.Buf[x] {
			// This can not error as we never exceed the bounds
			neighbours, _ := g.curr.AliveNeighbours(x, y)
			g.next.Buf[x][y] = NextState(neighbours, g.curr.Buf[x][y])
		}
	}

	g.curr, g.next = g.next, g.curr
}

func (g *Game) String() string {
	var sb strings.Builder

	for x := range g.curr.Buf {
		for y := range g.curr.Buf[x] {
			neighbours, _ := g.curr.AliveNeighbours(x, y)
			if neighbours > 1 {
				sb.WriteString(fmt.Sprintf("\033[0;3%dm", neighbours))
			}

			var s string
			if g.curr.Buf[x][y] {
				s = fmt.Sprintf("[%d]", neighbours)
			} else {
				s = " + "
			}
			sb.WriteString(s)

			sb.WriteString("\033[0m")
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
