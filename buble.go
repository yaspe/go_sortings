package main

import (
	"image/gif"
)

func drawBubleSort(array []int, anim *gif.GIF) {
	for a := 0; a < len(array); a++ {
		changed := false
		for b := 1; b < len(array)-a; b++ {
			if array[b-1] > array[b] {
				array[b-1], array[b] = array[b], array[b-1]
				drawArray(array, b-1, b, anim)
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}
