package main

import (
	"image"
	"image/color"
	"image/gif"

	"os"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 0, 255, 255}}

const (
	pointSize = 8
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

func swapAndDraw(array []int, s1, s2 int, anim *gif.GIF) {
	drawArray(array, s1, s2, anim)
	array[s1], array[s2] = array[s2], array[s1]
}

func drawArray(array []int, s1, s2 int, anim *gif.GIF) {
	rect := image.Rect(0, 0, canvSize, canvSize)
	img := image.NewPaletted(rect, palette)
	stepX := int(0.9*canvSize) / len(array)
	stepY := int(0.9*canvSize) / maxVal

	for x, y := range array {
		var ci uint8 = 1
		if x == s1 || x == s2 {
			ci = 2
		}
		drawPoint(int(0.1*canvSize)+x*stepX, int(0.1*canvSize)+y*stepY, img, ci)
	}

	anim.Delay = append(anim.Delay, 1)
	anim.Image = append(anim.Image, img)
}

func makeSortingGif(array []int, sorting func([]int, *gif.GIF), filename string) {
	anim := gif.GIF{}
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)
	drawArray(arrayCopy, -1, -1, &anim)

	sorting(arrayCopy, &anim)

	for i := 0; i < 100; i++ {
		drawArray(arrayCopy, -1, -1, &anim)
	}

	f, _ := os.Create(filename)
	defer f.Close()
	gif.EncodeAll(f, &anim)
}
