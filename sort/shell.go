package sort

import (
	"image/gif"

	"github.com/invzhi/sorting-visualization/animation"
)

// ShellSort will record a frame for every h-insertion sort.
func ShellSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	h := 1
	for h < l/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < l; i++ {
			t := a[i]
			j := i
			for ; j >= h && a[j-h] > t; j -= h {
				a[j] = a[j-h]
			}
			a[j] = t
		}
		h /= 3

		animation.SetLine(g, y, frame, a)
		frame++
	}
}
