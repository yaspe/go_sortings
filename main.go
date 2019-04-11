package main

import (
	"math/rand"
	"path"
)

func main() {
	const itemsNum = 100
	array := make([]int, itemsNum)
	for i := 0; i < itemsNum; i++ {
		array[i] = rand.Intn(maxVal)
	}

	steps1 := makeSortSteps(array, bubleSort)
	steps1.setPadding(0, 0)

	steps2 := makeSortSteps(array, insertionSort)
	steps2.setPadding(200, 0)

	steps3 := makeSortSteps(array, drawMergeSort)
	steps3.setPadding(0, 200)

	steps4 := makeSortSteps(array, quickSort)
	steps4.setPadding(200, 200)

	steps1.merge(&steps2).merge(&steps3).merge(&steps4)

	dir := "/Users/ya-spe/Downloads"
	makeSortingGifFromSteps(&steps1, path.Join(dir, "sortings.gif"))

}
