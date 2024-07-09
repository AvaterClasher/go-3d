package go3d

import (
	"image"
)

type Device struct {
	Width  int
	Height int

	ColorBuffer *image.RGBA
}

func NewDevice(width, height int) *Device {
	return &Device{
		Width:  width,
		Height: height,

		ColorBuffer: image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func (d *Device) Clear(c Color) {
	for y := 0; y < d.Height; y++ {
		for x := 0; x < d.Width; x++ {
			d.ColorBuffer.Set(x, y, c.NRGBA())
		}
	}
}

func (d *Device) Image() image.Image {
	return d.ColorBuffer
}

func (d *Device) SetPixel(x, y int, c Color) {
	d.ColorBuffer.Set(x, y, c.NRGBA())
}
