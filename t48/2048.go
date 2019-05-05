package t48

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
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

type Board [16]int

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
		b[adjI] *= 2
		b[i] = 0
	}

	if targetI != -1 {
		b[targetI] = b[i]
		b[i] = 0
	}
}
