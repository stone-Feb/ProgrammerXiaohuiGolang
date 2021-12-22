package chapter4

import (
	"ProgrammerXiaohuiGolang/edtion1/chapter4/part2"
	"ProgrammerXiaohuiGolang/edtion1/chapter4/part3"
	"ProgrammerXiaohuiGolang/edtion1/chapter4/part4"
	"fmt"
	"testing"
)

//记录有序边界的方式排序
func TestBubbleSort(t *testing.T)  {
	arr := []int{3,4,2,1,5,6,7,8}
	part2.BubbleSort(arr)
	fmt.Println(arr)
}

//鸡尾酒排序
func TestCockTailSort(t *testing.T)  {
	arr := []int{2,3,4,5,6,7,8,1}
	part2.CockTailSort(arr)
	fmt.Println(arr)
}

func TestStr(t *testing.T)  {
	str := "oJQRhxPwOAYkup8cAd-Vjg0htVJU"
	fmt.Println(len(str))
}

func TestQuickSort(t *testing.T)  {
	arr := []int{4,4,6,5,3,2,8,1}
	//arr := []int{4,7,6,5,3,2,8,1}
	//part3.QuickSort(arr , 0 , len(arr)-1)
	part3.QuickSortWithStack(arr , 0 , len(arr)-1)
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T)  {
	arr := []int{4,4,6,5,3,2,8,1}
	part4.HeapSort(arr)
	fmt.Println(arr)
}