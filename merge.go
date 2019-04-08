package main

import (
	"image/gif"
)

func merge(array []int, start, mid, end int, anim *gif.GIF) { // todo make inplace
	len1 := mid - start
	len2 := end - mid + 1

	left := make([]int, len1)
	copy(left, array[start:mid])

	right := make([]int, len2)
	copy(right, array[mid:end+1])

	cur1 := 0
	cur2 := 0
	for cur := start; cur <= end; cur++ {
		if cur1 < len1 && cur2 < len2 {
			if left[cur1] < right[cur2] {
				array[cur] = left[cur1]
				cur1++
			} else {
				array[cur] = right[cur2]
				cur2++
			}
		} else if cur1 < len1 {
			array[cur] = left[cur1]
			cur1++
		} else {
			array[cur] = right[cur2]
			cur2++
		}
		drawArray(array, cur, -1, anim)
	}
}

func drawMergeSortInd(array []int, start, end int, anim *gif.GIF) {
	if end-start == 0 { // 1 elem
		return
	} else if end-start == 1 {
		if array[start] > array[end] {
			swapAndDraw(array, start, end, anim)
		}
		return
	}

	mid := (start + end) / 2
	drawMergeSortInd(array, start, mid, anim)
	drawMergeSortInd(array, mid, end, anim)
	merge(array, start, mid, end, anim)
}

func drawMergeSort(array []int, anim *gif.GIF) {
	drawMergeSortInd(array, 0, len(array)-1, anim)
}
