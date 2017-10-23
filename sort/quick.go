package sort

import (
	"github.com/invzhi/sorting-visualization/gif256"
	"image/gif"
)

func partition(a []uint8) int {
	pivot, i := a[0], 1
	head, tail := 0, len(a)-1
	for i = 1; i <= tail; {
		if a[i] > pivot {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			// a[i], a[head] = a[head], a[i]
			a[i], a[i-1] = a[i-1], a[i]
			head++
			i++
		}
	}
	// return head
	return i - 1
}

func quickSort(a, data []uint8, y int, frame *int, g *gif.GIF) {
	if len(a) <= 1 {
		return
	}
	p := partition(a)

	gif256.SetLine(g, y, *frame, data)
	*frame = *frame + 1

	quickSort(a[:p], data, y, frame, g)
	quickSort(a[p+1:], data, y, frame, g)
}

func QuickSort(a []uint8, y int, g *gif.GIF) {
	data := a
	frame := 1
	quickSort(a, data, y, &frame, g)
}
