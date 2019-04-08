package main

import (
	"math/rand"
)

func main() {
	const itemsNum = 100
	array := make([]int, itemsNum)
	for i := 0; i < itemsNum; i++ {
		array[i] = rand.Intn(maxVal)
	}

	makeSortingGif(array, drawQuickSort, "/Users/ya-spe/Downloads/quick.gif")
	makeSortingGif(array, drawBubleSort, "/Users/ya-spe/Downloads/buble.gif")
	makeSortingGif(array, drawMergeSort, "/Users/ya-spe/Downloads/merge.gif")

}
