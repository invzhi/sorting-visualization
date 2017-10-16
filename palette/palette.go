package palette

import "image/color"

func hueToRGBA(h int) color.RGBA {
	var c color.RGBA

	hi := h / 60
	f := float64(h) / 60 - float64(hi)
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

func GetPalette(n int) color.Palette {
	colors := make(color.Palette, n)

	for h := 0; h < n; h++ {
		colors[h] = hueToRGBA(h % 360)
	}
	return colors
}
