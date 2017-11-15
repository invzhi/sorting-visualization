package sort

import (
	"image/gif"

	"github.com/invzhi/sorting-visualization/gif256"
)

func partition(a []uint8) int {
	l := len(a)
	pivot := a[0]
	i, j := 0, l

	for {
		for a[i+1] < pivot {
			i++
			if i+1 == l {
				break
			}
		}
		for ; a[j-1] > pivot; j-- {
		}
		i++
		j--
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[0], a[j] = a[j], a[0]
	return j
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
