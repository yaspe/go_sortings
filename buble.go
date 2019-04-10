package main

func bubleSort(array []int, steps *SortSteps) {
	for a := 0; a < len(array); a++ {
		changed := false
		for b := 1; b < len(array)-a; b++ {
			if array[b-1] > array[b] {
				steps.swapAndSave(array, b-1, b)
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}
