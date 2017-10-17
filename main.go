package main

import (
	"fmt"
	"github.com/invzhi/sorting-visualization/gif256"
	"github.com/invzhi/sorting-visualization/sort"
	"image/gif"
	"runtime"
	"sync"
)

const side = 256

func sortHue(sortF func([]uint8, int, *gif.GIF), fn string, delay int) {
	var wg sync.WaitGroup

	numCPU := runtime.NumCPU()
	sem := make(chan struct{}, numCPU)

	g, cis := gif256.NewRandGIF(side, side)

	fmt.Print(fn, " (0-255): ")
	wg.Add(side)
	for y := 0; y < side; y++ {
		sem <- struct{}{}
		go func(y int) {
			defer wg.Done()
			defer func() { <-sem }()
			sortF(cis[y], y, g)
		}(y)
		fmt.Print(y, ", ")
	}

	wg.Wait()
	gif256.EncodeGIF(g, fn, delay)
	fmt.Println("OK!")
}

func main() {
	sortHue(sort.SelectionSort, "gifs/selection.gif", 0)
	sortHue(sort.InsertionSort, "gifs/insertion.gif", 0)
	sortHue(sort.ShellSort, "gifs/shell.gif", 0)
	sortHue(sort.ShellSortOverview, "gifs/shell_overview.gif", 100)
}
