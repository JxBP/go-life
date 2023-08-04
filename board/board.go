package board

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Buf           [][]bool
	width, height int
}

func New(width int, height int) *Board {
	buf := make([][]bool, width)
	for i := 0; i < width; i++ {
		buf[i] = make([]bool, height)
	}
	return &Board{buf, width, height}
}

func (b *Board) Width() int {
	return b.width
}

func (b *Board) Height() int {
	return b.height
}

// Returns wether a set of coordinates is in the bounds of the board.
func (b *Board) IsInBounds(x int, y int) bool {
	return x >= 0 && x < b.width && y >= 0 && y < b.height
}

// Returns the number of alive neighbours.
// If the coordinates are out of bounds an error is returned.
// This function makes no promises about the number of alive neighbours returned if err is not nil.
func (b *Board) AliveNeighbours(x int, y int) (int, error) {
	if !b.IsInBounds(x, y) {
		return -1, fmt.Errorf("board: (x=%d, y=%d) out of bounds", x, y)
	}

	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if b.IsInBounds(x+i, y+j) && !(i == 0 && j == 0) && b.Buf[x+i][y+j] {
				neighbours++
			}
		}
	}
	return neighbours, nil
}

func (b *Board) String() string {
	var sb strings.Builder
	sb.WriteString("Board {\n")
	for row := range b.Buf {
		sb.WriteString("  { ")
		for col := range b.Buf[row] {
			sb.WriteString(strconv.Itoa(boolToInt(b.Buf[row][col])))

			// Only print a comma if this is NOT the last element
			if col < len(b.Buf[row])-1 {
				sb.WriteByte(',')
			}

			sb.WriteByte(' ')
		}
		sb.WriteString("}\n")
	}
	sb.WriteString("}")
	return sb.String()
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
