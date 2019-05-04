package t48

type Board [4 * 4]int

func (b Board) Left() Board {
	for i := 0; i < len(b); i++ {
		y := i / 4
		x := i % 4
		if x == 0 || b[i] == 0 {
			continue
		}
		target := -1
		for cursor := x - 1; cursor >= 0; cursor-- {
			v := y*4 + cursor
			if b[v] != 0 {
				break
			}
			target = v
		}

		if target == -1 {
			continue
		}

		b[target] = b[i]
		b[i] = 0

		_ = y
	}
	return b
}
