package sort

import (
	"image/gif"
	"github.com/invzhi/sorting-visualization/gif256"
)

func InsertionSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	e := 0
	for i := l - 1; i > 0; i-- {
		if a[i] < a[i - 1] {
			a[i], a[i - 1] = a[i - 1], a[i]
			e++
		}
	}
	if e == 0 {
		return
	}

	var j int
	for i := 1; i < l; i++ {
		t := a[i]
		for j = i; a[j - 1] > t; j-- {
			a[j] = a[j - 1]
		}
		a[j] = t

		gif256.SetLine(g, y, frame, a)
		frame++
	}
}