package board

import (
	"testing"
)

func logBoardSize(t *testing.T, b *Board) {
	t.Logf("Board size: width=%d height=%d", b.Width(), b.Height())
}

func logBoardState(t *testing.T, b *Board) {
	t.Logf("Board state:\n%s", b.String())
}

func TestIsInBoundsFalse(t *testing.T) {
	b := New(5, 5)
	logBoardSize(t, b)

	if b.IsInBounds(6, 5) {
		t.Error("IsInBounds(6, 5) = true, exepected: false")
	}

	if b.IsInBounds(5, 6) {
		t.Error("IsInBounds(5, 6) = true, exepected: false")
	}

	if b.IsInBounds(-1, 0) {
		t.Error("IsInBounds(-1, 0) = true, exepected: false")
	}

	if b.IsInBounds(0, -1) {
		t.Error("IsInBounds(0, -1) = true, exepected: false")
	}

	if b.IsInBounds(5, 5) {
		t.Error("IsInBounds(5, 5) = true, exepected: false")
	}
}

func TestIsInBoundsTrue(t *testing.T) {
	b := New(5, 5)
	logBoardSize(t, b)

	if !b.IsInBounds(0, 0) {
		t.Error("IsInBounds(0, 0) = false, expected: true")
	}

	if !b.IsInBounds(2, 4) {
		t.Error("IsInBounds(2, 4) = false, exepected: true")
	}
}

func TestAliveNeighbours(t *testing.T) {
	b := New(4, 4)
	logBoardSize(t, b)

	// This is the test board:
	//  (3) [2] (2)     (n) is a dead cell with n neighbours
	//  [3] [3] (2)     [n] is an alive cell with n neighbours
	//  [2] (3) (1)     The numbers indicate how many neighbours this cell has.

	b.Buf[0][1] = true
	b.Buf[1][0] = true
	b.Buf[1][1] = true
	b.Buf[2][0] = true

	logBoardState(t, b)

	type TestCase struct {
		x        int
		y        int
		expected int
	}
	inputs := []TestCase{
		{0, 0, 3}, {0, 1, 2}, {0, 2, 2},
		{1, 0, 3}, {1, 1, 3}, {1, 2, 2},
		{2, 0, 2}, {2, 1, 3}, {2, 2, 1},
	}

	for i := range inputs {
		tc := inputs[i]

		got, err := b.AliveNeighbours(tc.x, tc.y)
		if err != nil {
			t.Errorf("AliveNeighbours(%d, %d) failed: %v", tc.x, tc.y, err)
		} else if got != tc.expected {
			t.Errorf("AliveNeighbours(%d, %d) = %d, expected: %d", tc.x, tc.y, got, tc.expected)
		}
	}
}

func TestAliveNeighboursOoB(t *testing.T) {
	b := New(10, 10)
	logBoardSize(t, b)

	got, err := b.AliveNeighbours(10, 10)
	if err == nil {
		t.Errorf("AliveNeighbours(10, 10) = %d, expected: out of bounds error", got)
	}

	got, err = b.AliveNeighbours(9, 9)
	if err != nil {
		t.Errorf("AliveNeighbours(9, 9) failed: %v", err)
	} else if got != 0 {
		t.Errorf("AliveNeighbours(9, 9) = %d, expected: 0", got)
	}
}
