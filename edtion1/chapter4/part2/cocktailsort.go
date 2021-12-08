package part2

//冒泡排序：鸡尾酒排序（钟锤）

func CockTailSort(arr []int) {
	// 最外层len(arr)/2
	for i := 0; i < len(arr)/2; i++ {
		//奇数轮：钟锤从左往右走
		//有序标记
		isSorted := true
		for j := i; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
		//偶数轮：钟锤从右往左走
		isSorted = true
		for j := len(arr) - 1 - i; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}
