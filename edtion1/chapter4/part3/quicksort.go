package part3

//QuickSort 分治算法，快速排序
func QuickSort(arr []int, startIndex, endIndex int) {
	//递归退出条件
	if startIndex >= endIndex {
		return
	}
	//得到基准元素位置
	pivotIndex := partitionV2(arr, startIndex, endIndex)
	//根据基准元素位置，分成两部分递归排序
	//左递归
	QuickSort(arr, startIndex, pivotIndex-1)
	//右递归
	QuickSort(arr, pivotIndex+1, endIndex)
}

// 双边循环法：头尾双指针，头小尾大交换元素，指针重合pivot交换元素
func partition(arr []int, startIndex, endIndex int) int {
	// 第一个元素为基准元素
	pivot := arr[startIndex]
	left := startIndex
	right := endIndex

	for left != right {
		//尾指针左移动
		for left < right && arr[right] > pivot {
			right--
		}
		//头指针右移动
		for left < right && arr[left] <= pivot {
			left++
		}
		//头尾指针交换
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	//pivot和指针重合交换元素 left == right
	arr[startIndex], arr[left] = arr[left], pivot
	return left
}

//单边循环：单指针标记"小于基准元素的区域边界" + 循环遍历
func partitionV2(arr []int, startIndex, endIndex int) int {
	//基准元素
	pivot := arr[startIndex]
	//单指针标记
	mark := startIndex
	for i := startIndex ; i<= endIndex ; i++{
		//交换元素，标记后移
		if arr[i] < pivot {
			mark++
			arr[i] , arr[mark] = arr[mark] , arr[i]
		}
	}
	//循环结束，交换基准
	arr[startIndex] , arr[mark] = arr[mark] , pivot
	return mark
}
