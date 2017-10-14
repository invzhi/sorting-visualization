package rgbah

import "image/color"

type RGBAH struct {
	RGBA32 color.RGBA
	H int
}

func (c RGBAH) RGBA() (r, g, b, a uint32) {
	return c.RGBA32.RGBA()
}

func (c *RGBAH) generateRGBA() {
	hi := c.H / 60
	f := float64(c.H) / 60 - float64(hi)
	q := 1 - f

	switch hi {
	case 0:
		c.RGBA32.R = 0xFF
		c.RGBA32.G = uint8(f * 0xFF)
		c.RGBA32.B = 0x00
	case 1:
		c.RGBA32.R = uint8(q * 0xFF)
		c.RGBA32.G = 0xFF
		c.RGBA32.B = 0x00
	case 2:
		c.RGBA32.R = 0x00
		c.RGBA32.G = 0xFF
		c.RGBA32.B = uint8(f * 0xFF)
	case 3:
		c.RGBA32.R = 0x00
		c.RGBA32.G = uint8(q * 0xFF)
		c.RGBA32.B = 0xFF
	case 4:
		c.RGBA32.R = uint8(f * 0xFF)
		c.RGBA32.G = 0x00
		c.RGBA32.B = 0xFF
	case 5:
		c.RGBA32.R = 0xFF
		c.RGBA32.G = 0x00
		c.RGBA32.B = uint8(q * 0xFF)
	}
	c.RGBA32.A = 0xFF
}

func GetPalette(n int) color.Palette {
	// colors := make(color.Palette, 0, n)
	colors := make(color.Palette, n)

	for h := 0; h < n; h++ {
		var c RGBAH
		c.H = h % 360
		c.generateRGBA()
		// colors = append(colors, c)
		colors[h] = c
	}
	return colors
}
