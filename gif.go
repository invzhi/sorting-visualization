package main

import (
	"log"
	"sync"

	"github.com/invzhi/sorting-visualization/gif256"
)

func newGIF(sortf sorting, fn string, weight, height, delay int) {
	var wg sync.WaitGroup

	sem := make(chan struct{}, numCPU)

	g, cis := gif256.NewRandGIF(weight, height)

	wg.Add(height)
	log.Println(fn, "start generate")

	for y := 0; y < height; y++ {
		sem <- struct{}{}
		go func(y int) {
			defer wg.Done()
			defer func() { <-sem }()
			sortf(cis[y], y, g)
		}(y)
	}

	wg.Wait()
	gif256.EncodeGIF(g, fn, delay)
	log.Printf("%s generate %v frames\n", fn, len(g.Delay))
}
