package part1

import "fmt"

//第2章 数据结构基础
//2.1 数组：顺序表，顺序存储，随机读取
//动态数组代码参考 https://segmentfault.com/a/1190000015680429

type MyArray struct {
	//实际元素数量
	Size int
	//数组切片
	Array []int
}

// NewMyArray 构造函数
func NewMyArray(Size int) *MyArray {
	return &MyArray{
		Size:  0,
		Array: make([]int, Size),
	}
}

// Insert 数组插入元素：尾插入、中间插入、超限插入
// index 插入位置
// element 插入元素
func (a *MyArray) Insert(index, element int) {
	//判断访问下标是否超出范围
	if index < 0 || index > a.Size {
		panic("超出数组实际元素范围！")
	}
	//如果实际元素达到数组容量上限，数组扩容
	if a.Size >= len(a.Array) {
		a.reSize()
	}
	//将插入的索引位置之后的元素后移，腾出插入位置
	for i := a.Size - 1; i >= index; i-- {
		a.Array[i+1] = a.Array[i]
	}
	//插入新元素
	a.Array[index] = element
	a.Size++
}

// reSize 数组扩容
func (a *MyArray) reSize() {
	newArr := make([]int, a.Size*2)
	//不能使用append尾插入，需要头插入
	for i := 0; i < a.Size; i++ {
		newArr[i] = a.Array[i]
	}
	a.Array = newArr
}

// Delete 删除元素
func (a *MyArray) Delete(index int) (element int) {
	//判断访问下标是否超出范围
	if index < 0 || index >= a.Size {
		panic("超出数组实际元素范围！")
	}
	element = a.Array[index]
	//index 之后的元素，往左挪一位
	//TODO:原书错误：i < a.Size -1
	for i := index; i <= a.Size-1; i++ {
		a.Array[i] = a.Array[i+1]
	}
	a.Size--
	return
}

func (a *MyArray) Output() {
	for i := 0; i < a.Size; i++ {
		fmt.Println(a.Array[i])
	}
}
