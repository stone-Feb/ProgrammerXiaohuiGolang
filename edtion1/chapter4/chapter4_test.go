package chapter4

import (
	"ProgrammerXiaohuiGolang/edtion1/chapter4/part2"
	"fmt"
	"testing"
)

//记录有序边界的方式排序
func TestBubbleSort(t *testing.T)  {
	arr := []int{3,4,2,1,5,6,7,8}
	part2.BubbleSort(arr)
	fmt.Println(arr)
}