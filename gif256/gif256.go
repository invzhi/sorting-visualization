package gif256

import (
	"github.com/invzhi/sorting-visualization/palette"
	"image"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"sync"
)

var (
	pal = palette.GetPalette(256)
	m   sync.Mutex
)

func newFrame(g *gif.GIF) (img *image.Paletted) {
	w, h := g.Config.Width, g.Config.Height
	img = image.NewPaletted(image.Rect(0, 0, w, h), pal)
	// for x := 0; x < w; x++ { // maybe once
	// 	for y := 0; y < h; y++ {
	// 		img.SetColorIndex(x, y, )
	// 	}
	// }
	g.Image = append(g.Image, img)
	return
}

func NewRandGIF(h, w int) (*gif.GIF, [][]uint8) {
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
		for x, i := range rand.Perm(w) {
			cis[y][x] = uint8(i)
			// if i > 255, color index will overflow :P
			img.SetColorIndex(x, y, uint8(i))
		}
	}
	return g, cis
}

func SetLine(g *gif.GIF, y int, frame int, line []uint8) {
	var img *image.Paletted

	m.Lock()
	if frame >= len(g.Image) {
		img = newFrame(g)
	} else {
		img = g.Image[frame]
	}
	m.Unlock()

	for x, _ := range line {
		img.SetColorIndex(x, y, line[x])
	}
}

func setDelay(g *gif.GIF, delay int) {
	l := len(g.Image)
	g.Delay = make([]int, 0, l)

	for i := 0; i < l; i++ {
		g.Delay = append(g.Delay, delay)
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
