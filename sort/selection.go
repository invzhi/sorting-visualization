package sort

func SelectionSort(a []uint8, frame, set, done chan bool) {
	l := len(a)
	for i := 0; i < l; i++ {
		minIndex := i
		for j := i; j < l; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[minIndex], a[i] = a[i], a[minIndex]
		frame <- true
		<-set
	}
	done <- true
}
