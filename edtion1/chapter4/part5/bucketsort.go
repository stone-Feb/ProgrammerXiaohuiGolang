package part5

import (
	"container/heap"
	"fmt"
	"math"
)

//BucketSort 桶排序
func BucketSort(arr []float64) {
	//1. 得到数列的最大值和最小值，并算出差值
	max, min := arr[0], arr[0]
	for _, v := range arr {
		max = math.Max(max, v)
		min = math.Min(min, v)
	}
	d := max - min
	//2. 初始化桶
	bucketNum := len(arr)
	bucketList := make([]float64Heap, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucketList = append(bucketList, float64Heap{})
	}
	//3.遍历原是数组，将每个元素放入桶中
	for _, v := range arr {
		num := int((v - min) * (float64(bucketNum) - 1) / d)
		//bucketList[num].Push(v)
		//最小堆实现
		heap.Push(&bucketList[num],v)
	}
	//4. 对每个桶进行排序
	//	for i:=0 ; i< len(bucketList) ; i++{
	//		sort.Float64s(bucketList[i])
	//	}
	//5.输出全部元素
	sortedArray := make([]float64, bucketNum)
	index :=0
	for _, list := range bucketList {
		if list.Len() != 0 {
			for _,element := range list {
				sortedArray[index] = element
				index++
			}
		}
	}
	fmt.Println(sortedArray)
}

type float64Heap []float64

func (h float64Heap) Len() int           { return len(h) }
func (h float64Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h float64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *float64Heap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *float64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
