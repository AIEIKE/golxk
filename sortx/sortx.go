package sortx

func BubbleSort(a []int) {
	lastChangeIndex := 0
	sortBorder := len(a) - 1
	for i := 0; i < len(a)-1; i++ {
		isSorted := true
		for j := 0; j < sortBorder; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isSorted = false
				lastChangeIndex = j
			}
		}
		sortBorder = lastChangeIndex // 减少每轮比较个数
		if isSorted {
			break // 减少比较轮数
		}
	}
}

func CockTailSort(a []int) {
	var isSorted bool
	for i := 0; i < len(a)>>1; i++ {
		// 从左到右
		isSorted = true
		for j := i; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
		// 从右到左
		isSorted = true
		for j := len(a) - i - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}
