package sort

import (
	"image/gif"

	"github.com/invzhi/sorting-visualization/animation"
)

func sink(a []uint8, f, b int) {
	for 2*f+1 < b {
		c := 2*f + 1
		if c+1 < b && a[c] < a[c+1] {
			c++
		}
		if a[f] >= a[c] {
			break
		}
		a[f], a[c] = a[c], a[f]
		f = c
	}
}

// HeapSort will record a frame for every sink operation.
func HeapSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	for i := l/2 - 1; i >= 0; i-- {
		sink(a, i, l)

		animation.SetLine(g, y, frame, a)
		frame++
	}
	for i := l - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		sink(a, 0, i)

		animation.SetLine(g, y, frame, a)
		frame++
	}
}
