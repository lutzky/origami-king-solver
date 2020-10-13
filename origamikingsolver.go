package main

import (
	"fmt"
	"strings"
)

// Board represents the Origami King ring board
type Board struct {
	cells [4][12]int
	count int
	max   int
}

func (b *Board) String() string {
	var bs strings.Builder
	for _, r := range b.cells {
		for _, val := range r {
			ch := '.'
			if val != 0 {
				ch = rune('0' + val)
			}
			bs.WriteRune(ch)
		}
		bs.WriteRune('\n')
	}
	bs.WriteString(fmt.Sprintf("Count: %d\n", b.count))
	return bs.String()
}

// RingRotate rotates the specified ring by the given amount
func (b *Board) RingRotate(ring int, amount int) {
	var tmp [12]int
	for i := 0; i < 12; i++ {
		tmp[i] = b.cells[ring][(12+i-amount)%12]
	}
	for i := 0; i < 12; i++ {
		b.cells[ring][i] = tmp[i]
	}
}

// ColRotate rotates the specified column by the given amount
func (b *Board) ColRotate(col int, amount int) {
	var tmp [8]int
	for i := 0; i < 4; i++ {
		tmp[(8+i+amount)%8] = b.cells[i][col]
		tmp[(8+4+i+amount)%8] = b.cells[3-i][(col+6)%12]
	}
	for i := 0; i < 4; i++ {
		b.cells[i][col] = tmp[(8+i)%8]
		b.cells[3-i][(col+6)%12] = tmp[(8+4+i)%8]
	}
}

// ParseBoard parses a board from s. A '.' represents no enemy, '0' through '9'
// represent enemy classes.
func ParseBoard(s string) Board {
	var result Board
	for i, c := range strings.Join(strings.Fields(s), "") {
		val := 0
		row := i / 12
		col := i % 12
		if c != '.' {
			val = int(c - '0')
			result.count++
		}
		if val > result.max {
			result.max = val
		}

		result.cells[row][col] = val
	}
	return result
}

// IsVictory returns true iff every enemy is either in a full column with the
// same type as itself, or in a 2x2 box in rows 0-1 with the same type as
// itself.
func (b *Board) IsVictory() bool {
	for col := 0; col < 12; col++ {
		if b.cells[0][col] != 0 &&
			b.cells[0][col] == b.cells[1][col] &&
			b.cells[0][col] == b.cells[2][col] &&
			b.cells[0][col] == b.cells[3][col] {
			b.cells[0][col] = 0
			b.cells[1][col] = 0
			b.cells[2][col] = 0
			b.cells[3][col] = 0
			b.count -= 4
		}
	}

	for col := 0; col < 12; col++ {
		ncol := (col + 1) % 12
		if b.cells[0][col] != 0 &&
			b.cells[0][col] == b.cells[1][col] &&
			b.cells[0][col] == b.cells[0][ncol] &&
			b.cells[0][col] == b.cells[1][ncol] {
			b.cells[0][col] = 0
			b.cells[1][col] = 0
			b.cells[0][ncol] = 0
			b.cells[1][ncol] = 0
			b.count -= 4
		}
	}

	return b.count == 0
}

// Solve returns the set of steps to solve the board, or {"FAIL"}
// if it cannot be solved in the given number of moves.
func (b Board) Solve(moves int) []string {
	if moves == 0 {
		if b.IsVictory() {
			return []string{"OK"}
		}
		return []string{"FAIL"}
	}

	for ring := 0; ring < 4; ring++ {
		for amount := 1; amount < 13; amount++ {
			b.RingRotate(ring, 1)
			tested := b.Solve(moves - 1)
			if tested[len(tested)-1] == "OK" {
				if amount > 6 {
					amount -= 12
				}
				return append([]string{fmt.Sprintf("RING(%d, %d)", ring, amount)}, tested...)
			}
		}
	}

	for col := 0; col < 12; col++ {
		for amount := 1; amount < 5; amount++ {
			b.ColRotate(col, 1)
			if col == 2 {
			}
			tested := b.Solve(moves - 1)
			if tested[len(tested)-1] == "OK" {
				if amount > 6 {
					amount -= 12
				}
				return append([]string{fmt.Sprintf("COL(%d, %d)", col, amount)}, tested...)
			}
		}
	}

	return []string{"FAIL"}
}

func main() {
	fmt.Println("vim-go")
}
