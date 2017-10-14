package main

import (
	"os"
	"fmt"
	"log"
	// "time"
	"image"
	"image/gif"
	"math/rand"
	"github.com/invzhi/sorting-visualization/rgbah"
	"github.com/invzhi/sorting-visualization/sort"
)

func main() {
	const side = 256
	var (
		images []*image.Paletted
		delays []int
	)
	var datas [side]sort.Data
	palette := rgbah.GetPalette(side)

	img := image.NewPaletted(image.Rect(0, 0, side, side), palette)
	// shuffle
	for y := 0; y < side; y++ {
		datas[y].Pixels = make([]rgbah.RGBAH, side)
		for x, i := range rand.Perm(side) {
			datas[y].Pixels[x] = palette[i].(rgbah.RGBAH)
			img.Set(x, y, palette[i])
		}
	}
	images = append(images, img)
	delays = append(delays, 0)

	leftSide := side

	frame := make(chan bool)
	done := make(chan bool)
	set := make(chan bool)

	// sort
	for y := 0; y < side; y++ {
		go datas[y].SelectionSort(frame, set, done)
	}
	
	// set
	go func() {
		for a := 0; ; a++ {
			img := image.NewPaletted(image.Rect(0, 0, side, side), palette)
			// time.Sleep(time.Second)
			for i := 0; i < leftSide; i++ {
				<-frame
			}
			fmt.Println(a, "frame")

			for y := 0; y < side; y++ {
				for x := 0; x < side; x++ {
					img.Set(x, y, datas[y].Pixels[x])
				}
			}

			for i := 0; i < leftSide; i++ {
				set <- true
			}

			images = append(images, img)
			delays = append(delays, 0)
		}
	}()

	for leftSide > 0 {
		<-done
		fmt.Println("sorted.")
		leftSide--
	}

	f, err := os.Create("gifs/image.gif")
	if err != nil {
		log.Fatal(err)
	}

	g := &gif.GIF{
		Image: images,
		Delay: delays,
	}

	if err := gif.EncodeAll(f, g); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
