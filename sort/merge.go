package sort

import (
	"github.com/invzhi/sorting-visualization/gif256"
	"image/gif"
)

func quickMerge(a, t []uint8, lo, mid, hi int) {
	for i := 0; i <= mid; i++ {
		t[i] = a[i]
	}
	for i := hi; i > mid; i-- {
		t[i] = a[mid+1+hi-i]
	}

	i, j := lo, hi
	for k := lo; k <= hi; k++ {
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
	ll := len(a)

	t := make([]uint8, ll)
	// bottom up
	for l := 1; l < ll; l += l {
		for lo := 0; lo < ll-l; lo += l + l {
			mid := lo + l - 1
			hi := mid + l
			if ll-1 < hi {
				hi = ll - 1
			}
			quickMerge(a, t, lo, mid, hi)
		}
		gif256.SetLine(g, y, frame, a)
		frame++
	}
}
