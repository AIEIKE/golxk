package main

import (
	"fmt"
	"golxk/sortx"
)

func main() {
	a := []int{3, 4, 2, 1, 5, 6, 7, 8}
	sortx.QuickSort(a, 0, len(a)-1)
	fmt.Println("排序", a)
}
