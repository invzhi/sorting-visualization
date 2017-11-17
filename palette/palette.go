package palette

import "image/color"

const round = 360

var pal = make(color.Palette, round)

func init() {
	for hue := 0; hue < round; hue++ {
		pal[hue] = hueToRGBA(hue)
	}
}

// GetPalette return palette, hue is between 0 to n-1.
func GetPalette(n int) color.Palette {
	var colors color.Palette
	if n > round {
		colors = pal[:]
	} else {
		colors = pal[:n]
	}
	return colors
}

func hueToRGBA(h int) color.RGBA {
	var c color.RGBA

	hi := h / 60
	f := float64(h)/60 - float64(hi)
	q := 1 - f

	switch hi {
	case 0:
		c.R = 0xFF
		c.G = uint8(f * 0xFF)
		c.B = 0x00
	case 1:
		c.R = uint8(q * 0xFF)
		c.G = 0xFF
		c.B = 0x00
	case 2:
		c.R = 0x00
		c.G = 0xFF
		c.B = uint8(f * 0xFF)
	case 3:
		c.R = 0x00
		c.G = uint8(q * 0xFF)
		c.B = 0xFF
	case 4:
		c.R = uint8(f * 0xFF)
		c.G = 0x00
		c.B = 0xFF
	case 5:
		c.R = 0xFF
		c.G = 0x00
		c.B = uint8(q * 0xFF)
	}
	c.A = 0xFF
	return c
}
