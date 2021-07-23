package main

import (
	"fmt"
	"golxk/sortx"
)

func main() {
	a := []int{3, 4, 2, 1, 5, 6, 7, 8}
	sortx.CockTailSort(a)
	fmt.Println("排序", a)
}
