package main

func merge(array []int, start, mid, end int, steps *SortSteps) { // todo make inplace
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
		steps.append(array, cur, -1)
	}
}

func drawMergeSortInd(array []int, start, end int, steps *SortSteps) {
	if end-start == 0 { // 1 elem
		return
	} else if end-start == 1 {
		if array[start] > array[end] {
			steps.swapAndSave(array, start, end)
		}
		return
	}

	mid := (start + end) / 2
	drawMergeSortInd(array, start, mid, steps)
	drawMergeSortInd(array, mid, end, steps)
	merge(array, start, mid, end, steps)
}

func drawMergeSort(array []int, steps *SortSteps) {
	drawMergeSortInd(array, 0, len(array)-1, steps)
}
