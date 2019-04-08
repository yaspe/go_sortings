package main

import (
	"image/gif"
)

func drawQuickSortInd(array []int, left, right int, anim *gif.GIF) {
	if left >= right {
		return
	}

	pivotpt := (left + right) / 2
	pivot := array[pivotpt]

	i := left
	j := right - 1

	swapAndDraw(array, pivotpt, right, anim)
	pivotpt = right

	for i < j {
		for i < right-1 && array[i] < pivot {
			i++
		}

		for j > 0 && array[j] > pivot {
			j--
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
			drawArray(array, i, j, anim)
			i++
			j--
		}
	}

	if i == j && array[i] < array[pivotpt] {
		i++
	}

	swapAndDraw(array, i, pivotpt, anim)

	drawQuickSortInd(array, left, i-1, anim)
	drawQuickSortInd(array, i+1, right, anim)
}

func drawQuickSort(array []int, anim *gif.GIF) {
	drawQuickSortInd(array, 0, len(array)-1, anim)
}
