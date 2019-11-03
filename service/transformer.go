package service

import (
	"image"
	"image/draw"
	"image/jpeg"
	"math"
	"os"
)

type Grid struct {
	Height int
	Width  int
}

type Transformer struct {
	Grid   Grid
	Canvas *image.RGBA
	Count  int
	Length int
	Size   int
}

func NewCanvas(size int, length int) *Transformer {
	canvas := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: size * length,
			Y: size * length,
		},
	})
	return &Transformer{Canvas: canvas, Length: length, Size: size}
}

func (t *Transformer) Merge(img image.Image) {
	id := t.Count

	x := id % t.Size
	y := math.Round(float64(id / t.Size))

	minPoint := image.Point{x * t.Length, int(y) * t.Length}
	maxPoint := minPoint.Add(image.Point{t.Length, t.Length})
	nextRect := image.Rectangle{minPoint, maxPoint}

	draw.Draw(t.Canvas, nextRect, img, image.ZP, draw.Src)
	t.Count++
}

func (t *Transformer) Export() {
	out, err := os.Create("./merged.jpg")
	if err != nil {
	}

	var opt jpeg.Options
	opt.Quality = 80
	jpeg.Encode(out, t.Canvas, &opt)
}
