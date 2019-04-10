package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 0, 255, 255}}

const (
	pointSize = 5
	canvSize  = 400
	maxVal    = 100
)

func drawPoint(x, y int, img *image.Paletted, ci uint8) {
	for i := x - pointSize/2; i < x+pointSize/2; i++ {
		for j := y - pointSize/2; j < y+pointSize/2; j++ {
			img.SetColorIndex(i, j, ci)
		}
	}
}

func makeSortingGifFromSteps(s *SortSteps, filename string) {
	anim := gif.GIF{}

	for _, step := range s.steps {
		drawStep(&step, &anim)
	}

	f, _ := os.Create(filename)
	defer f.Close()
	gif.EncodeAll(f, &anim)
}

func drawStep(s *SortStep, anim *gif.GIF) {
	rect := image.Rect(0, 0, canvSize, canvSize)
	img := image.NewPaletted(rect, palette)
	stepX := 1.25 //int(0.35*canvSize) / len(s.array)
	stepY := 1.25 //int(0.35*canvSize) / maxVal

	for i, point := range s.array {
		x := point.x
		y := point.y
		var ci uint8 = 1
		if x == s.p1 || x == s.p2 {
			ci = 2
		}
		drawPoint(s.x[i]+int(0.1*canvSize+float64(x)*stepX), s.y[i]+int(0.1*canvSize+float64(y)*stepY), img, ci)
	}

	anim.Delay = append(anim.Delay, 1)
	anim.Image = append(anim.Image, img)
}

func makeSortSteps(array []int, sorting func([]int, *SortSteps)) SortSteps {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	steps := SortSteps{}
	steps.append(array, -1, -1)

	sorting(arrayCopy, &steps)
	return steps
}
