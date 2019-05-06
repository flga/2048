package t48

import (
	"reflect"
	"testing"
)

func TestMoveLeft(t *testing.T) {
	tests := []struct {
		name      string
		board     Board
		want      Board
		wantMoves []Move
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{0, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{0, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Merge,
				},
				Move{
					From: Point{2, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{0, 0},
					Type: Merge,
				},
				Move{
					From: Point{3, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Merge,
				},
				Move{
					From: Point{2, 1},
					To:   Point{1, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{0, 1},
					Type: Merge,
				},
				Move{
					From: Point{3, 1},
					To:   Point{1, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Merge,
				},
				Move{
					From: Point{2, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{0, 3},
					Type: Merge,
				},
				Move{
					From: Point{3, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Merge,
				},
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{0, 3},
					Type: Slide,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, moves := tt.board.Move(Left, nil)
			if !tt.want.cmp(got) {
				t.Fatalf("Move(Left) want %v, got %v", tt.want.pretty(), got.pretty())
			}
			if !reflect.DeepEqual(tt.wantMoves, moves) {
				t.Fatalf("Move(Left) want moves %v, got %v", tt.wantMoves, moves)
			}
		})
	}
}

func TestMoveRight(t *testing.T) {
	tests := []struct {
		name      string
		board     Board
		want      Board
		wantMoves []Move
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{2, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{3, 1},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Merge,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Merge,
				},
				Move{
					From: Point{0, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{3, 0},
					Type: Merge,
				},
				Move{
					From: Point{1, 0},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Merge,
				},
				Move{
					From: Point{0, 1},
					To:   Point{2, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{3, 1},
					Type: Merge,
				},
				Move{
					From: Point{1, 1},
					To:   Point{2, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Merge,
				},
				Move{
					From: Point{0, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{3, 3},
					Type: Merge,
				},
				Move{
					From: Point{1, 3},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Merge,
				},
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, moves := tt.board.Move(Right, nil)
			if !tt.want.cmp(got) {
				t.Fatalf("Move(Right) want %v, got %v", tt.want.pretty(), got.pretty())
			}
			if !reflect.DeepEqual(tt.wantMoves, moves) {
				t.Fatalf("Move(Left) want moves %v, got %v", tt.wantMoves, moves)
			}
		})
	}
}

func TestMoveUp(t *testing.T) {
	tests := []struct {
		name      string
		board     Board
		want      Board
		wantMoves []Move
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 3},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 1},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 1},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 1},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 1},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 2},
					Type: Slide,
				},
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
				2, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: nil,
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
				0, 2, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: nil,
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
				2, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{2, 0},
					Type: Slide,
				},
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
				0, 2, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
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
				2, 0, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
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
				0, 2, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
				2, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: nil,
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
				0, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: nil,
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
				0, 0, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: nil,
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
				2, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
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
				0, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{2, 0},
					Type: Slide,
				},
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
				0, 0, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
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
				2, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
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
				0, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
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
				0, 0, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
				2, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: nil,
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
				0, 2, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: nil,
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
				2, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{2, 0},
					Type: Slide,
				},
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
				0, 2, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
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
				2, 2, 2, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
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
				0, 2, 2, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{2, 3},
					To:   Point{2, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Slide,
				},
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
				2, 4, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Merge,
				},
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
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Merge,
				},
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
				0, 4, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Merge,
				},
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
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Merge,
				},
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
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Merge,
				},
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
				0, 4, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Merge,
				},
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
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Merge,
				},
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
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Merge,
				},
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
				0, 4, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Merge,
				},
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
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Merge,
				},
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
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Merge,
				},
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
				0, 4, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Merge,
				},
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
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Merge,
				},
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
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 3},
					To:   Point{0, 0},
					Type: Merge,
				},
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
				0, 4, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 0},
					Type: Merge,
				},
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
				0, 0, 0, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 0},
					Type: Merge,
				},
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
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Merge,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 1},
					Type: Slide,
				},
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
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Merge,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 1},
					Type: Slide,
				},
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
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Merge,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 1},
					Type: Slide,
				},
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
				4, 0, 0, 0,
				2, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 0},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 0},
					Type: Merge,
				},
				Move{
					From: Point{0, 3},
					To:   Point{0, 1},
					Type: Slide,
				},
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
				0, 4, 0, 0,
				0, 2, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 0},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 0},
					Type: Merge,
				},
				Move{
					From: Point{1, 3},
					To:   Point{1, 1},
					Type: Slide,
				},
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
				0, 0, 0, 4,
				0, 0, 0, 2,
				0, 0, 0, 0,
				0, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 0},
					Type: Slide,
				},
				Move{
					From: Point{3, 2},
					To:   Point{3, 0},
					Type: Merge,
				},
				Move{
					From: Point{3, 3},
					To:   Point{3, 1},
					Type: Slide,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, moves := tt.board.Move(Up, nil)
			if !tt.want.cmp(got) {
				t.Fatalf("Move(Up) want %v, got %v", tt.want.pretty(), got.pretty())
			}
			if !reflect.DeepEqual(tt.wantMoves, moves) {
				t.Fatalf("Move(Left) want moves %v, got %v", tt.wantMoves, moves)
			}
		})
	}
}

func TestMoveDown(t *testing.T) {
	tests := []struct {
		name      string
		board     Board
		want      Board
		wantMoves []Move
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{2, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{2, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 2},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{2, 2},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{2, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{0, 0},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{2, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{2, 2},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 2},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 2},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 2},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 1},
					Type: Slide,
				},
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
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 2},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{3, 1},
					Type: Slide,
				},
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
			wantMoves: nil,
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
			wantMoves: nil,
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 2, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 2,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 2, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 2,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Slide,
				},
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
				2, 0, 2, 0,
			},
			wantMoves: nil,
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
				0, 2, 0, 2,
			},
			wantMoves: nil,
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 2, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 2, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 2,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 2, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 2, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 2, 2,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{2, 3},
					Type: Slide,
				},
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
				2, 2, 0, 0,
			},
			wantMoves: nil,
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
				0, 2, 2, 0,
			},
			wantMoves: nil,
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
				0, 0, 2, 2,
			},
			wantMoves: nil,
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 2, 2, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 2, 2,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 0},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 2, 2, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{2, 1},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 2, 2,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{2, 1},
					To:   Point{2, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Slide,
				},
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
				2, 2, 2, 0,
			},
			wantMoves: nil,
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
				0, 2, 2, 2,
			},
			wantMoves: nil,
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 4, 0, 2,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 0},
					To:   Point{3, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 4, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				4, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 4, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 4,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Merge,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Slide,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 3},
					Type: Merge,
				},
				Move{
					From: Point{0, 0},
					To:   Point{0, 2},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 4, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Slide,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 3},
					Type: Merge,
				},
				Move{
					From: Point{1, 0},
					To:   Point{1, 2},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Slide,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 3},
					Type: Merge,
				},
				Move{
					From: Point{3, 0},
					To:   Point{3, 2},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{0, 2},
					To:   Point{0, 3},
					Type: Merge,
				},
				Move{
					From: Point{0, 1},
					To:   Point{0, 2},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 2, 0, 0,
				0, 4, 0, 0,
			},
			wantMoves: []Move{
				Move{
					From: Point{1, 2},
					To:   Point{1, 3},
					Type: Merge,
				},
				Move{
					From: Point{1, 1},
					To:   Point{1, 2},
					Type: Slide,
				},
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
				0, 0, 0, 0,
				0, 0, 0, 2,
				0, 0, 0, 4,
			},
			wantMoves: []Move{
				Move{
					From: Point{3, 2},
					To:   Point{3, 3},
					Type: Merge,
				},
				Move{
					From: Point{3, 1},
					To:   Point{3, 2},
					Type: Slide,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, moves := tt.board.Move(Down, nil)
			if !tt.want.cmp(got) {
				t.Fatalf("Move(Down) want %v, got %v", tt.want.pretty(), got.pretty())
			}
			if !reflect.DeepEqual(tt.wantMoves, moves) {
				t.Fatalf("Move(Left) want moves %v, got %v", tt.wantMoves, moves)
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

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		d    []Move
		want []Move
	}{
		{name: "nil", d: nil, want: nil},
		{name: "empty", d: []Move{}, want: []Move{}},
		{
			name: "single",
			d: []Move{
				Move{Type: 1},
			},
			want: []Move{
				Move{Type: 1},
			},
		},
		{
			name: "even",
			d: []Move{
				Move{Type: 1},
				Move{Type: 0},
			},
			want: []Move{
				Move{Type: 0},
				Move{Type: 1},
			},
		},
		{
			name: "odd",
			d: []Move{
				Move{Type: 1},
				Move{Type: 1},
				Move{Type: 0},
			},
			want: []Move{
				Move{Type: 0},
				Move{Type: 1},
				Move{Type: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.d)
			if !reflect.DeepEqual(tt.d, tt.want) {
				t.Fatalf("reverse() want %v, got %v", tt.want, tt.d)
			}
		})
	}
}
