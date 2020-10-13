package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRingRotate(t *testing.T) {
	testCases := []struct {
		name             string
		in               Board
		ring, ringAmount int
		col, colAmount   int
		want             Board
	}{
		{
			name: "ring 1 -3",
			in: ParseBoard(`
............
.1.....1....
...22..1....
...22..1....
			`),
			ring: 1, ringAmount: -3,
			want: ParseBoard(`
............
....1.....1.
...22..1....
...22..1....
`),
		},
		{
			name: "col 1 1",
			in: ParseBoard(`
............
.1.....1....
...22..1....
...22..1....
			`),
			col: 1, colAmount: -1,
			want: ParseBoard(`
.1..........
............
...22..1....
.1.22..1....
`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("In:\n%s", tc.in.String())
			got := tc.in
			got.RingRotate(tc.ring, tc.ringAmount)
			t.Logf("RingRotate(%d, %d) = \n%s", tc.ring, tc.ringAmount, got.String())
			got.ColRotate(tc.col, tc.colAmount)
			t.Logf("ColRotate(%d, %d) = \n%s", tc.col, tc.colAmount, got.String())
			if got.String() != tc.want.String() {
				t.Errorf("Got:\n%s\n; want:\n%s", got.String(), tc.want.String())
			}
		})
	}
}

func TestIsVictory(t *testing.T) {
	testCases := []struct {
		name string
		in   Board
		want bool
	}{
		{
			name: "random",
			in: ParseBoard(`
.1..........
............
...22..1....
.1.22..1....
`),
			want: false,
		},
		{
			name: "onebox",
			in: ParseBoard(`
...22.......
...22.......
............
............
`),
			want: true,
		},
		{
			name: "onecol",
			in: ParseBoard(`
...1........
...1........
...1........
...1........
`),
			want: true,
		},
		{
			name: "both",
			in: ParseBoard(`
...1.22.....
...1.22.....
...1........
...1........
`),
			want: true,
		},
		{
			name: "confusing",
			in: ParseBoard(`
...11.......
...11.......
...1........
...1........
`),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.in.IsVictory()
			if got != tc.want {
				t.Errorf("got: %t; want: %t", got, tc.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	testCases := []struct {
		name  string
		in    Board
		moves int
		want  []string
	}{
		{
			name: "trivial",
			in: ParseBoard(`
						...1........
						...1........
						...1........
						...1........
						`),
			moves: 0,
			want:  []string{"OK"},
		},
		{
			name: "easy",
			in: ParseBoard(`
						...1........
						...1........
						......1.....
						...1........
						`),
			moves: 1,
			want:  []string{"RING(2, -3)", "OK"},
		},
		{
			name: "ring-2",
			in: ParseBoard(`
						...1........
						....1.......
						......1.....
						...1........
						`),
			moves: 2,
			want:  []string{"RING(1, -1)", "RING(2, -3)", "OK"},
		},
		{
			name: "ring-3",
			in: ParseBoard(`
						...1........
						....1.......
						......1.....
						..1.........
						`),
			moves: 3,
			want:  []string{"RING(0, 1)", "RING(2, -2)", "RING(3, 2)", "OK"},
		},
		{
			name: "easy-col",
			in: ParseBoard(`
			...1.....1..
			...1.....1..
			............
			............
			`),
			moves: 1,
			want:  []string{"COL(3, 2)", "OK"},
		},
		{
			name: "easy-col2",
			in: ParseBoard(`
...1....1...
...1....1...
............
............
`),
			moves: 1,
			want:  []string{"COL(2, 2)", "OK"},
		},
		{
			name: "super-annoying-hard",
			in: ParseBoard(`
...22..1....
...22..1....
.1.....1....
............
`),
			moves: 3,
			want:  []string{"RING(1, 6)", "COL(1, 1)", "RING(1, 6)", "OK"},
		},
	}

	for _, tc := range testCases {
		got := tc.in.Solve(tc.moves)
		if d := cmp.Diff(tc.want, got); d != "" {
			t.Errorf("diff (-want +got):\n%s", d)
		}
	}
}
