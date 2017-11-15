package sort

import (
	"image/gif"

	"github.com/invzhi/sorting-visualization/animation"
)

func RadixSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	const BASE = 10
	b := make([]uint8, l)
	max := int(a[0])
	exp := 1

	for i := 1; i < l; i++ {
		if int(a[i]) > max {
			max = int(a[i])
		}
	}

	for max/exp > 0 {
		bucket := make([]int, BASE)
		for i := 0; i < l; i++ {
			radix := int(a[i]) / exp % BASE
			bucket[radix]++
		}
		for i := 1; i < BASE; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := l - 1; i >= 0; i-- {
			radix := int(a[i]) / exp % BASE
			bucket[radix]--
			b[bucket[radix]] = a[i]
		}
		for i := 0; i < l; i++ {
			a[i] = b[i]
		}
		exp *= BASE

		animation.SetLine(g, y, frame, a)
		frame++
	}
}
