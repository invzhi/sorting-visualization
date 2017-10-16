package main

import (
	"fmt"
	"sync"
	"runtime"
	"github.com/invzhi/sorting-visualization/gif256"
	"github.com/invzhi/sorting-visualization/sort"
)

const side = 256

func main() {
	var wg sync.WaitGroup

	numCPU := runtime.NumCPU()
	sem := make(chan struct{}, numCPU)

	g, cis := gif256.NewRandGIF(side, side)

	wg.Add(side)
	for y := 0; y < side; y++ {
		sem <- struct{}{}
		go func(y int) {
			defer wg.Done()
			defer func() {<-sem}()
			sort.SelectionSort(cis[y], y, g)
		}(y)
		fmt.Print(y, " ")
	}

	wg.Wait()
	gif256.EncodeGIF(g, "gifs/se.gif")
	fmt.Println("Good")
}
