package part2

//冒泡排序：记录有序数列的边界

func BubbleSort(arr []int) {
	//最后一次交换的位置
	lastExchangeIndex := 0
	//无序数列的边界，每次比较只需要比较到这里
	sortBorder := len(arr) - 1
	//最外层循环：控制所有回合
	for i := 0; i < len(arr); i++ {
		//有序标记：每一轮初始都是true
		isSorted := true
		//内部循环实现每一轮的冒泡处理，进行元素交换
		for j := 0; j < sortBorder; j++ {
			//比较交换
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//有元素交换，不是有序
				isSorted = false
				//把无序数列的边界更新为最后一次交换元素的位置
				lastExchangeIndex = j
			}
		}
		sortBorder = lastExchangeIndex
		if isSorted {
			break
		}
	}
}
