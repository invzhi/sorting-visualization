package sort

import (
	"github.com/invzhi/sorting-visualization/gif256"
	"image/gif"
)

func BubbleSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	for i := 0; i < l - 1; i++ {
		for j := l-1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
		gif256.SetLine(g, y, frame, a)
		frame++
	}
}
