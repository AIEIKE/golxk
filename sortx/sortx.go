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
