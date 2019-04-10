package main

func quickSortInd(array []int, left, right int, steps *SortSteps) {
	if left >= right {
		return
	}

	pivotpt := (left + right) / 2
	pivot := array[pivotpt]

	i := left
	j := right - 1

	steps.swapAndSave(array, pivotpt, right)
	pivotpt = right

	for i < j {
		for i < right-1 && array[i] < pivot {
			i++
		}

		for j > 0 && array[j] > pivot {
			j--
		}

		if i < j {
			steps.swapAndSave(array, i, j)
			i++
			j--
		}
	}

	if i == j && array[i] < array[pivotpt] {
		i++
	}

	steps.swapAndSave(array, i, pivotpt)

	quickSortInd(array, left, i-1, steps)
	quickSortInd(array, i+1, right, steps)
}

func quickSort(array []int, steps *SortSteps) {
	quickSortInd(array, 0, len(array)-1, steps)
}
