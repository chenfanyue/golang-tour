package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 25, 350, 60)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{255, 190, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
