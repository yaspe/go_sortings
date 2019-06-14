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

	c1 := make(chan SortSteps)
	go makeSortSteps(array, bubleSort, c1)

	c2 := make(chan SortSteps)
	go makeSortSteps(array, insertionSort, c2)

	c3 := make(chan SortSteps)
	go makeSortSteps(array, drawMergeSort, c3)

	c4 := make(chan SortSteps)
	go makeSortSteps(array, quickSort, c4)

	steps1 := <-c1
	steps1.setPadding(0, 0)

	steps2 := <-c2
	steps2.setPadding(200, 0)

	steps3 := <-c3
	steps3.setPadding(0, 200)

	steps4 := <-c4
	steps4.setPadding(200, 200)

	steps1.merge(&steps2).merge(&steps3).merge(&steps4)

	dir := "/Users/ya-spe/Downloads"
	makeSortingGifFromSteps(&steps1, path.Join(dir, "sortings.gif"))

}
