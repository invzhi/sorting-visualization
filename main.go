package main

import (
	"flag"
	"image/gif"
	"log"
	"runtime"
	"strings"

	"github.com/invzhi/sorting-visualization/sort"
)

type sorting func([]uint8, int, *gif.GIF)

const (
	minWeight = 5
)

var (
	numCPU int

	sortName, filename    string
	weight, height, delay int

	m = map[string]sorting{
		"selection": sort.SelectionSort,
		"insertion": sort.InsertionSort,
		"shell":     sort.ShellSort,
		"merge":     sort.MergeSort,
		"quick":     sort.QuickSort,
		"heap":      sort.HeapSort,
		"bubble":    sort.BubbleSort,
		"radix":     sort.RadixSort,
	}
)

func init() {
	numCPU = runtime.NumCPU()

	const (
		defaultWeight = 256
		defaultHeight = 256
		defaultDelay  = 10

		sortUsage  = "selection, insertion, shell, merge, quick, bubble, radix, all"
		delayUsage = "successive delay times, one per frame, in 100ths of a second"
	)
	flag.StringVar(&sortName, "sorting", "", sortUsage)
	flag.StringVar(&filename, "filename", "", "GIF's filename (default: sorting name)")

	flag.IntVar(&weight, "weight", defaultWeight, "GIF's weight")
	flag.IntVar(&height, "height", defaultHeight, "GIF's height")
	flag.IntVar(&delay, "delay", defaultDelay, delayUsage)
}

func main() {
	flag.Parse()

	sortf, isexist := m[sortName]
	if isexist == false && sortName != "all" {
		log.Fatalln("sorting is not existed:", sortName)
	}

	// filename
	if filename == "" {
		filename = sortName
	}
	if strings.HasSuffix(filename, ".gif") == false {
		filename += ".gif"
	}

	// weight, height
	if weight < minWeight {
		log.Fatalln("weight can not less than", minWeight)
	}
	if height <= 0 {
		log.Fatalln("height can not less than 1")
	}

	// delay
	if delay < 0 {
		log.Fatalln("delay can not less than 0")
	}

	// if sorting name is "all", generate all GIF
	if sortName != "all" {
		newGIF(sortf, filename, weight, height, delay)
	} else {
		// no matter about filename
		for name, sortf := range m {
			newGIF(sortf, name+".gif", weight, height, delay)
		}
	}
}
