package part5

import "fmt"

//计数排序: 利用数组下标来确定元素的正确位置。元素之间不比较

//CountSort 数值较小，差值不大
func CountSort(arr []int) []int {
	//1 获取数列的最大值
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	//2. 根据数列的最大值确定统计数组的长度
	max++
	countArray := make([]int, max)
	//3. 遍历数列，填充数组
	for _, v := range arr {
		countArray[v]++
	}
	//4. 遍历统计数组，输出结果
	index := 0
	sortedArray := make([]int, len(arr))
	for k, v := range countArray {
		//数列中没有的不输出
		if v == 0 {
			continue
		}
		for i := 1; i <= v; i++ {
			sortedArray[index] = k
			index++
		}
	}
	return sortedArray
}

//CountSortV2 数值较大，稳定排序
func CountSortV2(arr []int) []int{
	//1. 得到数列的最大值和最小值，并算出差值
	max := arr[0]
	min := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	d := max - min
	// 2. 创建统计数组并统计对应元素个数
	countArray := make([]int, d+1)
	for _, v := range arr {
		key := v - min
		countArray[key]++
	}
	//3. 统计数组做变形，后面的元素等于前面元素之和
	for k, _ := range countArray {
		//第二个元素开始，每一个元素都加上前面所有元素之和
		if k ==0 {
			continue
		}
		countArray[k] += countArray[k-1]
	}
	//4. 倒序遍历原始数列，从统计数组中找到正确位置，输出到结果数组
	sortedArray := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		s := arr[i]-min
		if countArray[s] == 0 {
			continue
		}
		skey := countArray[s]-1
		sortedArray[skey] = arr[i]
		//统计数组元素-1，方便下个原始数组做稳定排序
		countArray[s]--
	}
	fmt.Println(sortedArray)
	return sortedArray
}
