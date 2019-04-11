package main

func insertionSort(array []int, steps *SortSteps) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			steps.swapAndSave(array, j, j-1)
		}
	}
}
