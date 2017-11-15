package sort

import (
	"image/gif"

	"github.com/invzhi/sorting-visualization/gif256"
)

func InsertionSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	for i := 1; i < l; i++ {
		t := a[i]
		j := i
		for ; j > 0 && a[j-1] > t; j-- {
			a[j] = a[j-1]
		}
		a[j] = t

		gif256.SetLine(g, y, frame, a)
		frame++
	}
}
