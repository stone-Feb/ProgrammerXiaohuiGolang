package part4

import "fmt"

//堆排序：无序数列先调整为最小堆，删除推顶，下沉调整
//顺序存储方式的二叉树：leftChild = 2 * parent + 1 rightChild = 2 * parent + 2

func HeapSort(arr []int) {
	//1. 无序数组构建成最大堆
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		downAdjust(arr, i, len(arr))
	}
	fmt.Println("first :", arr)
	//2.循环交换集合尾部元素到堆顶，并调节堆产生新的堆顶
	for i := len(arr) - 1; i > 0; i-- {
		//堆顶出堆
		arr[i], arr[0] = arr[0], arr[i]
		//下沉调整最大堆
		downAdjust(arr, 0, i)
	}
}

/**
 * 下沉调整
 * @param array     待调整的堆
 * @param parentIndex    要下沉的父节点
 * @param length    堆的有效大小
 */
func downAdjust(arr []int, parentIndex, length int) {
	//temp 保存父节点值，用于最后下沉附值
	temp := arr[parentIndex]
	//默认是左孩子节点
	childIndex := 2*parentIndex + 1
	for childIndex < length {
		//如果有右孩子 && 右孩子大于左孩子节点的值，则定位到右孩子
		if childIndex+1 < length && arr[childIndex+1] > arr[childIndex] {
			childIndex++
		}
		//如果父节点>=任一孩子节点，直接跳出
		if temp >= arr[childIndex] {
			break
		}
		//交换值
		arr[parentIndex] = arr[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	//父节点下沉
	arr[parentIndex] = temp
}
