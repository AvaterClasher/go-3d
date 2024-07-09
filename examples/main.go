package main

import (
	"image/png"
	"image"
	"os"
	"github.com/AvaterClasher/go-3d/go-3d"
)

func main() {
	d := go3d.NewDevice(256, 256)
	d.Clear(go3d.BLACK)
	d.SetPixel(128,128,go3d.WHITE)
	d.SetPixel(120,120,go3d.WHITE)
	d.SetPixel(110,110,go3d.WHITE)
	d.SetPixel(100,100,go3d.WHITE)
	Save("examples/example.png", d.Image())
}

func Save(filename string, img image.Image) {
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}
}