package sort

import (
	"github.com/invzhi/sorting-visualization/rgbah"
)

type Data struct {
	Pixels []rgbah.RGBAH
}

func (d *Data) SelectionSort(frame, set, done chan bool) {
	side := len(d.Pixels)
	for i := 0; i < side; i++ {
		minIndex := i
		for j := i; j < side; j++ {
			if d.Pixels[j].H < d.Pixels[minIndex].H {
				minIndex = j
			}
		}
		d.Pixels[minIndex], d.Pixels[i] = d.Pixels[i], d.Pixels[minIndex]
		frame <- true
		<-set
	}
	done <- true
}
