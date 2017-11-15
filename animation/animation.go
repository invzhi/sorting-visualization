package animation

import (
	"image"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/invzhi/sorting-visualization/palette"
)

var (
	pal = palette.GetPalette(256)
	m   sync.Mutex
)

func newFrame(g *gif.GIF) *image.Paletted {
	w, h := g.Config.Width, g.Config.Height
	r := image.Rect(0, 0, w, h)
	pix := make([]uint8, w*h)

	for i := range pix {
		pix[i] = uint8(i % w)
	}
	img := &image.Paletted{pix, 1 * w, r, pal}
	g.Image = append(g.Image, img)
	return img
}

func NewRandGIF(w, h int) (*gif.GIF, [][]uint8) {
	g := &gif.GIF{
		Image: make([]*image.Paletted, 0),
		Config: image.Config{
			ColorModel: pal,
			Width:      w,
			Height:     h,
		},
	}
	img := newFrame(g)
	cis := make([][]uint8, h)

	for y := 0; y < h; y++ {
		cis[y] = make([]uint8, w)
		// if ci > 255, color index will overflow :P
		for x, ci := range rand.Perm(w) {
			cis[y][x] = uint8(ci)
			img.Pix[y*w+x] = uint8(ci)
		}
	}
	return g, cis
}

func SetLine(g *gif.GIF, y, frame int, line []uint8) {
	var img *image.Paletted

	m.Lock()
	if frame >= len(g.Image) {
		img = newFrame(g)
	} else {
		img = g.Image[frame]
	}
	m.Unlock()

	// left begin pixel
	l := img.PixOffset(0, y)
	for x, ci := range line {
		img.Pix[l+x] = ci
	}
}

func setDelay(g *gif.GIF, delay int) {
	g.Delay = make([]int, len(g.Image))

	for i := range g.Image {
		g.Delay[i] = delay
	}
}

func EncodeGIF(g *gif.GIF, fn string, delay int) {
	f, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}

	setDelay(g, delay)

	if err := gif.EncodeAll(f, g); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
