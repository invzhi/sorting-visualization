package sort

import (
	"github.com/invzhi/sorting-visualization/gif256"
	"image/gif"
)

func quickMerge(a, t []uint8, mid int) {
	l := len(a)

	for i := 0; i <= mid; i++ {
		t[i] = a[i]
	}
	for i := l - 1; i > mid; i-- {
		t[i] = a[mid+l-i] // mid+l-i = mid+1 + l-1-i
	}

	i, j := 0, l-1
	for k := 0; k < l; k++ {
		if t[i] < t[j] {
			a[k] = t[i]
			i++
		} else {
			a[k] = t[j]
			j--
		}
	}
}

func MergeSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	t := make([]uint8, l)
	// bottom up
	for blk := 1; blk < l; blk += blk {
		for lo := 0; lo < l-blk; lo += blk + blk {
			mid := lo + blk - 1
			hi := mid + blk
			if l-1 < hi {
				hi = l - 1
			}
			quickMerge(a[lo:hi+1], t[lo:hi+1], mid-lo)
		}

		gif256.SetLine(g, y, frame, a)
		frame++
	}
}
