package t48

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func (d Direction) cursor(i int) (x, y int) {
	switch d {
	case Left:
		return i%4 - 1, i / 4
	case Right:
		return i%4 + 1, i / 4
	case Up:
		return i % 4, i/4 - 1
	case Down:
		return i % 4, i/4 + 1
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
	return true
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

type Board [4 * 4]int

func (b Board) Move(d Direction) Board {
	if d == Left || d == Up {
		for i := 0; i < len(b); i++ {
			process(&b, d, i)
		}
	}

	if d == Right || d == Down {
		for i := len(b) - 1; i >= 0; i-- {
			process(&b, d, i)
		}
	}

	return b
}

func process(b *Board, d Direction, i int) {
	if b[i] == 0 {
		return
	}

	var (
		targetX   = i % 4
		targetY   = i / 4
		targetIdx = -1
	)
	for x, y := d.cursor(i); d.limit(x, y); x, y, _ = d.inc(x, y) {
		v := y*4 + x
		if b[v] != 0 {
			break
		}
		targetX = x
		targetY = y
		targetIdx = v
	}

	adjX, adjY, adjInBounds := d.inc(targetX, targetY)
	adjIdx := adjY*4 + adjX
	if adjInBounds && b[adjIdx] == b[i] {
		b[adjIdx] *= 2
		b[i] = 0
	}
	if targetIdx != -1 {
		b[targetIdx] = b[i]
		b[i] = 0
	}
}
