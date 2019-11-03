package service

import (
	"image"
	"image/color"
	"math/rand"
	"testing"
	"time"
)

func TestTransformer_Merge(t *testing.T) {
	size := 3
	gridLength := 300

	img1 := GenerateRandomGrid(gridLength)
	img2 := GenerateRandomGrid(gridLength)
	img3 := GenerateRandomGrid(gridLength)
	img4 := GenerateRandomGrid(gridLength)
	img5 := GenerateRandomGrid(gridLength)
	img6 := GenerateRandomGrid(gridLength)
	img7 := GenerateRandomGrid(gridLength)
	img8 := GenerateRandomGrid(gridLength)
	img9 := GenerateRandomGrid(gridLength)

	transformer := NewCanvas(size, gridLength)

	transformer.Merge(img1)
	transformer.Merge(img2)
	transformer.Merge(img3)
	transformer.Merge(img4)
	transformer.Merge(img5)
	transformer.Merge(img6)
	transformer.Merge(img7)
	transformer.Merge(img8)
	transformer.Merge(img9)

	transformer.Export()

	t.Error("How might this fail?")
}

func TestTransformer_NewCanvas(t *testing.T) {
	size := 5
	gridLength := 350
	transformer := NewCanvas(size, gridLength)
	expectedSize := size * gridLength

	t.Log("bounds:", transformer.Canvas.Bounds().Dx())
	got := transformer.Canvas.Bounds().Dx()
	if got != expectedSize {
		t.Errorf("Canvas is not the expected size, got %d want %d", got, expectedSize)
	}
}

func GenerateRandomGrid(length int) *image.RGBA {
	rand.Seed(time.Now().Unix())
	img := image.NewRGBA(image.Rect(0, 0, length, length))
	rand.Seed(time.Now().UnixNano())
	color := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 0xff}

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			img.Set(x, y, color)
		}
	}

	return img
}
