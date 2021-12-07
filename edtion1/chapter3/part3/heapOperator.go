package part3

//最小二叉堆操作
//二叉堆是顺序存储方式，所以节点都保存在数组中。父节点下标是 p 左孩子：2*p+1 右孩子 2*p+2

// UpAdjust 插入节点
func UpAdjust(arr []int) {
	childIndex := len(arr) - 1
	parentIndex := (childIndex - 1) / 2
	//最后的叶子节点暂存，最后附值给空出来的下标
	temp := arr[childIndex]
	for childIndex > 0 && temp < arr[parentIndex] {
		//单向附值
		arr[childIndex] = arr[parentIndex]
		//交换下标
		childIndex, parentIndex = parentIndex, (parentIndex-1)/2
	}
	//附值给最后一次交换下标后，空出来的parentIndex
	arr[childIndex] = temp
}

// DownAdjust 删除节点
func DownAdjust(arr []int, parentIndex int) {
	temp := arr[parentIndex]
	childIndex := 2*parentIndex + 1
	length := len(arr)
	for childIndex < length {
		//二叉堆：父节点大于或小于所有孩子节点
		//存在右孩子 && 右孩子值 > 左孩子值。确定要比较的子节点
		if childIndex+1 < length && arr[childIndex+1] < arr[childIndex] {
			childIndex++
		}
		//父节点小于子节点，就不用交换
		if temp <= arr[childIndex] {
			break
		}
		arr[parentIndex] = arr[childIndex]
		parentIndex, childIndex = childIndex, childIndex*2+1
	}
	arr[parentIndex] = temp
}

// BuildHeap 无序二叉树构建二叉堆
func BuildHeap(arr []int) {
	//从最后一个非叶子节点开始，依次下沉调整
	for i := len(arr) / 2; i >= 0; i-- {
		DownAdjust(arr , i)
	}
}
