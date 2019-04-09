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

	dir := "/Users/ya-spe/Downloads"
	makeSortingGif(array, drawQuickSort, path.Join(dir, "quick.gif"))
	makeSortingGif(array, drawBubleSort, path.Join(dir, "buble.gif"))
	makeSortingGif(array, drawMergeSort, path.Join(dir, "merge.gif"))

}
