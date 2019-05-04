package t48

import (
	"testing"
)

func TestMoveLeft(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  Board
	}{
		{
			name: "empty board",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (2)",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
				0, 0, 0, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.Left()
			if !tt.want.cmp(got) {
				t.Fatalf("Left() want %v, got %v", tt.want.pretty(), got.pretty())
			}
		})
	}
}

func TestMoveRight(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  Board
	}{
		{
			name: "empty board",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 4,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (2)",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
				0, 0, 0, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.Right()
			if !tt.want.cmp(got) {
				t.Fatalf("Right() want %v, got %v", tt.want.pretty(), got.pretty())
			}
		})
	}
}

func (b Board) pretty() [][]int {
	return [][]int{
		b[0:4],
		b[4:8],
		b[8:12],
		b[12:16],
	}
}

func (b Board) cmp(a Board) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
