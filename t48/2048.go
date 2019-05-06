package t48

type MoveType int

const (
	Slide MoveType = iota
	Merge
)

type Point struct {
	X, Y int
}

type Move struct {
	From, To Point
	Type     MoveType
}

// Direction that the pieces can be moved
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) cursor(x, y int) (int, int) {
	switch d {
	case Left:
		return x - 1, y
	case Right:
		return x + 1, y
	case Up:
		return x, y - 1
	case Down:
		return x, y + 1
	}
	return -1, -1
}

func (d Direction) limit(x, y int) bool {
	switch d {
	case Left:
		return x >= 0
	case Right:
		return x < 4
	case Up:
		return y >= 0
	case Down:
		return y < 4
	}
	return false
}

func (d Direction) inc(x, y int) (int, int, bool) {
	switch d {
	case Left:
		x--
	case Right:
		x++
	case Up:
		y--
	case Down:
		y++
	}
	return x, y, x > -1 && x < 4 && y > -1 && y < 4
}

// Board stores the immutable state of a game.
// Methods of board that need to modify state will return a copy instead.
type Board [16]int

// Move shifts all the pieces in the given direction.
//
// Each piece will be moved to the furthest available spot.
// If a piece lands in a cell whose adjacent value is the same, they are merged
// into the adjacent cell and its value gets doubled.
func (b Board) Move(d Direction, moves []Move) (Board, []Move) {
	if d == Left || d == Up {
		for i := 0; i < len(b); i++ {
			moves = process(&b, d, i, moves)
		}
	}

	if d == Right || d == Down {
		for i := len(b) - 1; i >= 0; i-- {
			moves = process(&b, d, i, moves)
		}
		// reverse(moves)
	}

	return b, moves
}

func process(b *Board, d Direction, i int, moves []Move) []Move {
	if b[i] == 0 {
		return moves
	}

	var (
		curX    = i % 4
		curY    = i / 4
		targetX = curX
		targetY = curY
		targetI = -1
	)
	for x, y := d.cursor(curX, curY); d.limit(x, y); x, y, _ = d.inc(x, y) {
		target := y*4 + x
		if b[target] != 0 {
			break
		}
		targetX = x
		targetY = y
		targetI = target
	}

	adjX, adjY, adjInBounds := d.inc(targetX, targetY)
	adjI := adjY*4 + adjX
	if adjInBounds && b[adjI] == b[i] {
		moves = append(moves, Move{
			From: Point{curX, curY},
			To:   Point{adjX, adjY},
			Type: Merge,
		})
		b[adjI] *= 2
		b[i] = 0
	} else if targetI != -1 {
		moves = append(moves, Move{
			From: Point{curX, curY},
			To:   Point{targetX, targetY},
			Type: Slide,
		})
		b[targetI] = b[i]
		b[i] = 0
	}

	return moves
}

func reverse(m []Move) {
	length := len(m)

	for i := 0; i < length/2; i++ {
		a := i
		b := length - 1 - i
		m[a], m[b] = m[b], m[a]
	}
}
