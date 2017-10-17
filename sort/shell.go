package sort

import (
	"github.com/invzhi/sorting-visualization/gif256"
	"image/gif"
)

func ShellSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	h := 1
	for h < l/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		var j int
		for i := h; i < l; i++ {
			t := a[i]
			for j = i; j >= h && a[j-h] > t; j -= h {
				a[j] = a[j-h]
			}
			a[j] = t

			gif256.SetLine(g, y, frame, a)
			frame++
		}
		h /= 3
	}
}
