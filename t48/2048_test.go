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
			name: "single item top - lhs",
			board: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item top - middle",
			board: Board{
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item middle - lhs",
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
			name: "single item middle - middle",
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
			name: "single item middle - rhs",
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
			name: "single item bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "single item bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "single item bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - lhs",
			board: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - middle",
			board: Board{
				0, 4, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - rhs",
			board: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
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
			name: "two consecutive items (H) middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
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
			name: "two consecutive items (H) middle - rhs",
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
			name: "two consecutive items (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - lhs",
			board: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - middle",
			board: Board{
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - rhs",
			board: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) top - lhs",
			board: Board{
				2, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) top - rhs",
			board: Board{
				0, 2, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
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
			name: "two items with gap (H) middle - rhs",
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
			name: "two items with gap (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 0,
			},
		},
		{
			name: "two items with gap (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 0,
			},
		},
		{
			name: "two items with gap (V) top - lhs",
			board: Board{
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) top - rhs",
			board: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) top - lhs",
			board: Board{
				2, 4, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) top - rhs",
			board: Board{
				2, 0, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) middle - lhs",
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
			name: "three items with gap (H) middle - rhs",
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
			name: "three items with gap (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
		},
		{
			name: "three items with gap (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
		},
		{
			name: "three items with gap (V) top - lhs",
			board: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				8, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) top - middle",
			board: Board{
				0, 0, 2, 0,
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 8, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 8,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) bottom - lhs",
			board: Board{
				8, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				8, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) bottom - middle",
			board: Board{
				0, 0, 8, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
			want: Board{
				8, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) bottom - rhs",
			board: Board{
				0, 0, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
			want: Board{
				8, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) top - lhs",
			board: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) top - rhs",
			board: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) middle - lhs",
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
			name: "three consecutive items (H) middle - rhs",
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
		{
			name: "three consecutive items (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
		},
		{
			name: "three consecutive items (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
		},
		{
			name: "three consecutive items (V) top - lhs",
			board: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap top - lhs",
			board: Board{
				2, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap top - rhs",
			board: Board{
				0, 2, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap middle - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive top - lhs",
			board: Board{
				2, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive top - middle",
			board: Board{
				0, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive top - rhs",
			board: Board{
				0, 0, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive middle - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple top - lhs",
			board: Board{
				2, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple top - rhs",
			board: Board{
				0, 2, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple middle - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 2, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 2, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
		},
		{
			name: "merge (H) multiple bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 2, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
		},
		{
			name: "merge (H) edges",
			board: Board{
				0, 2, 0, 0,
				2, 0, 0, 2,
				0, 0, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap top - lhs",
			board: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive top - lhs",
			board: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive middle - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple top - lhs",
			board: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "merge (V) edges",
			board: Board{
				0, 2, 0, 0,
				2, 0, 0, 2,
				0, 0, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.Move(Left)
			if !tt.want.cmp(got) {
				t.Fatalf("Move(Left) want %v, got %v", tt.want.pretty(), got.pretty())
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
			name: "single item top - lhs",
			board: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item top - middle",
			board: Board{
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item middle - lhs",
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
			name: "single item middle - middle",
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
			name: "single item middle - rhs",
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
			name: "single item bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "single item bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "single item bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "two consecutive items (H) top - lhs",
			board: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - middle",
			board: Board{
				0, 4, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - rhs",
			board: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
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
			name: "two consecutive items (H) middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
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
			name: "two consecutive items (H) middle - rhs",
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
			name: "two consecutive items (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
		},
		{
			name: "two consecutive items (H) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
		},
		{
			name: "two consecutive items (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
		},
		{
			name: "two consecutive items (V) top - lhs",
			board: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - middle",
			board: Board{
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - rhs",
			board: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "two consecutive items (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "two consecutive items (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "two items with gap (H) top - lhs",
			board: Board{
				2, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 2, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) top - rhs",
			board: Board{
				0, 2, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 2, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
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
			name: "two items with gap (H) middle - rhs",
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
			name: "two items with gap (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 4,
			},
		},
		{
			name: "two items with gap (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 4,
			},
		},
		{
			name: "two items with gap (V) top - lhs",
			board: Board{
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) top - rhs",
			board: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
		},
		{
			name: "two items with gap (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
		},
		{
			name: "two items with gap (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
		},
		{
			name: "three items with gap (H) top - lhs",
			board: Board{
				2, 4, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) top - rhs",
			board: Board{
				2, 0, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) middle - lhs",
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
			name: "three items with gap (H) middle - rhs",
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
			name: "three items with gap (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
		},
		{
			name: "three items with gap (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
		},
		{
			name: "three items with gap (V) top - lhs",
			board: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				8, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 8,
			},
		},
		{
			name: "three items with gap (V) top - middle",
			board: Board{
				0, 0, 2, 0,
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 8, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 8,
			},
		},
		{
			name: "three items with gap (V) top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 8,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 8,
			},
		},
		{
			name: "three items with gap (V) bottom - lhs",
			board: Board{
				8, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "three items with gap (V) bottom - middle",
			board: Board{
				0, 0, 8, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "three items with gap (V) bottom - rhs",
			board: Board{
				0, 0, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "three consecutive items (H) top - lhs",
			board: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) top - rhs",
			board: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) middle - lhs",
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
			name: "three consecutive items (H) middle - rhs",
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
		{
			name: "three consecutive items (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
		},
		{
			name: "three consecutive items (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
		},
		{
			name: "three consecutive items (V) top - lhs",
			board: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
		},
		{
			name: "three consecutive items (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
		},
		{
			name: "three consecutive items (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
		},
		{
			name: "merge (H) gap top - lhs",
			board: Board{
				2, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap top - rhs",
			board: Board{
				0, 2, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap middle - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) gap bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
		},
		{
			name: "merge (H) gap bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
		},
		{
			name: "merge (H) consecutive top - lhs",
			board: Board{
				2, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive top - middle",
			board: Board{
				0, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive top - rhs",
			board: Board{
				0, 0, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive middle - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) consecutive bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
		},
		{
			name: "merge (H) consecutive bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
		},
		{
			name: "merge (H) consecutive bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
		},
		{
			name: "merge (H) multiple top - lhs",
			board: Board{
				2, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 2, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple top - rhs",
			board: Board{
				0, 2, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 2, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple middle - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 2, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 2, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 2, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (H) multiple bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 2, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 4,
			},
		},
		{
			name: "merge (H) multiple bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 2, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 4,
			},
		},
		{
			name: "merge (H) edges",
			board: Board{
				0, 2, 0, 0,
				2, 0, 0, 2,
				0, 0, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) gap top - lhs",
			board: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) gap bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) gap bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) gap bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) consecutive top - lhs",
			board: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive middle - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) consecutive bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) consecutive bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) consecutive bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) multiple top - lhs",
			board: Board{
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "merge (V) multiple bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) multiple bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) multiple bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 2,
				0, 0, 0, 2,
			},
		},
		{
			name: "merge (V) edges",
			board: Board{
				0, 2, 0, 0,
				2, 0, 0, 2,
				0, 0, 0, 0,
				0, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.Move(Right)
			if !tt.want.cmp(got) {
				t.Fatalf("Move(Right) want %v, got %v", tt.want.pretty(), got.pretty())
			}
		})
	}
}

func TestMoveUp(t *testing.T) {
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
			name: "single item top - lhs",
			board: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item top - middle",
			board: Board{
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "single item bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - lhs",
			board: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - middle",
			board: Board{
				0, 4, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 4, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - rhs",
			board: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 4, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
			want: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
			},
			want: Board{
				0, 4, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
			want: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - lhs",
			board: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - middle",
			board: Board{
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - rhs",
			board: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 4, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) top - lhs",
			board: Board{
				2, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) top - rhs",
			board: Board{
				0, 2, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
			},
			want: Board{
				2, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
			},
			want: Board{
				0, 2, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) top - lhs",
			board: Board{
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 4, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) top - rhs",
			board: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
			},
			want: Board{
				0, 0, 2, 0,
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) top - lhs",
			board: Board{
				2, 4, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) top - rhs",
			board: Board{
				2, 0, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
			},
			want: Board{
				2, 4, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
			},
			want: Board{
				2, 0, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) top - lhs",
			board: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				8, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) top - middle",
			board: Board{
				0, 0, 2, 0,
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 8, 0,
			},
			want: Board{
				0, 0, 2, 0,
				0, 0, 4, 0,
				0, 0, 8, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 8,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) bottom - lhs",
			board: Board{
				8, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				8, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) bottom - middle",
			board: Board{
				0, 0, 8, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 8, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) bottom - rhs",
			board: Board{
				0, 0, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 8,
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) top - lhs",
			board: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) top - rhs",
			board: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
			want: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
			want: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - lhs",
			board: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
			want: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
			},
			want: Board{
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
			want: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.Move(Up)
			if !tt.want.cmp(got) {
				t.Fatalf("Move(Up) want %v, got %v", tt.want.pretty(), got.pretty())
			}
		})
	}
}

func TestMoveDown(t *testing.T) {
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
			name: "single item top - lhs",
			board: Board{
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "single item top - middle",
			board: Board{
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
			},
		},
		{
			name: "single item top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "single item middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "single item middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
			},
		},
		{
			name: "single item middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "single item bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "single item bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
			},
		},
		{
			name: "single item bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
			},
		},
		{
			name: "two consecutive items (H) top - lhs",
			board: Board{
				4, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) top - middle",
			board: Board{
				0, 4, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
			},
		},
		{
			name: "two consecutive items (H) top - rhs",
			board: Board{
				0, 0, 4, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
		},
		{
			name: "two consecutive items (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
			},
		},
		{
			name: "two consecutive items (H) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
		},
		{
			name: "two consecutive items (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 2, 0, 0,
			},
		},
		{
			name: "two consecutive items (H) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 2, 0,
			},
		},
		{
			name: "two consecutive items (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 2,
			},
		},
		{
			name: "two consecutive items (V) top - lhs",
			board: Board{
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - middle",
			board: Board{
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 0, 0,
				0, 2, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) top - rhs",
			board: Board{
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "two consecutive items (V) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - middle",
			board: Board{
				0, 0, 0, 0,
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 0, 0,
				0, 2, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "two consecutive items (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
		},
		{
			name: "two consecutive items (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "two items with gap (H) top - lhs",
			board: Board{
				2, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
			},
		},
		{
			name: "two items with gap (H) top - rhs",
			board: Board{
				0, 2, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
			},
		},
		{
			name: "two items with gap (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
			},
		},
		{
			name: "two items with gap (H) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
			},
		},
		{
			name: "two items with gap (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 0,
			},
		},
		{
			name: "two items with gap (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 4,
			},
		},
		{
			name: "two items with gap (V) top - lhs",
			board: Board{
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
		},
		{
			name: "two items with gap (V) top - rhs",
			board: Board{
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "two items with gap (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 4, 0,
			},
		},
		{
			name: "two items with gap (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
			},
		},
		{
			name: "three items with gap (H) top - lhs",
			board: Board{
				2, 4, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
			},
		},
		{
			name: "three items with gap (H) top - rhs",
			board: Board{
				2, 0, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
			},
		},
		{
			name: "three items with gap (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
			},
		},
		{
			name: "three items with gap (H) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
			},
		},
		{
			name: "three items with gap (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 8,
			},
		},
		{
			name: "three items with gap (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 4, 8,
			},
		},
		{
			name: "three items with gap (V) top - lhs",
			board: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				8, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) top - middle",
			board: Board{
				0, 0, 2, 0,
				0, 0, 4, 0,
				0, 0, 0, 0,
				0, 0, 8, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 2, 0,
				0, 0, 4, 0,
				0, 0, 8, 0,
			},
		},
		{
			name: "three items with gap (V) top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
		},
		{
			name: "three items with gap (V) bottom - lhs",
			board: Board{
				8, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				8, 0, 0, 0,
				4, 0, 0, 0,
				2, 0, 0, 0,
			},
		},
		{
			name: "three items with gap (V) bottom - middle",
			board: Board{
				0, 0, 8, 0,
				0, 0, 0, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 8, 0,
				0, 0, 4, 0,
				0, 0, 2, 0,
			},
		},
		{
			name: "three items with gap (V) bottom - rhs",
			board: Board{
				0, 0, 0, 8,
				0, 0, 0, 0,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 8,
				0, 0, 0, 4,
				0, 0, 0, 2,
			},
		},
		{
			name: "three consecutive items (H) top - lhs",
			board: Board{
				2, 4, 8, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
		},
		{
			name: "three consecutive items (H) top - rhs",
			board: Board{
				0, 2, 4, 8,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
		},
		{
			name: "three consecutive items (H) middle - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
		},
		{
			name: "three consecutive items (H) middle - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
		},
		{
			name: "three consecutive items (H) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 8, 0,
			},
		},
		{
			name: "three consecutive items (H) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 4, 8,
			},
		},
		{
			name: "three consecutive items (V) top - lhs",
			board: Board{
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - middle",
			board: Board{
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) top - rhs",
			board: Board{
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
				0, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
		},
		{
			name: "three consecutive items (V) bottom - lhs",
			board: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				8, 0, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - middle",
			board: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
			},
			want: Board{
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 4, 0, 0,
				0, 8, 0, 0,
			},
		},
		{
			name: "three consecutive items (V) bottom - rhs",
			board: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
			want: Board{
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
				0, 0, 0, 8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.Move(Down)
			if !tt.want.cmp(got) {
				t.Fatalf("Move(Down) want %v, got %v", tt.want.pretty(), got.pretty())
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
