package part1

import "fmt"

//第2章 数据结构基础
//2.1 数组：顺序表，顺序存储，随机读取
//动态数组代码参考 https://segmentfault.com/a/1190000015680429

type MyArray struct {
	//实际元素数量
	size int
	//数组切片
	array []int
}

// NewMyArray 构造函数
func NewMyArray(size int) *MyArray {
	return &MyArray{
		size:  0,
		array: make([]int, size),
	}
}

// Insert 数组插入元素
// index 插入位置
// element 插入元素
func (a *MyArray) Insert(index, element int) {
	//判断访问下标是否超出范围
	if index < 0 || index > a.size {
		panic("超出数组实际元素范围！")
	}
	//如果实际元素达到数组容量上线，数组扩容
	if a.size >= len(a.array) {
		a.resize()
	}
	//将插入的索引位置之后的元素后移，腾出插入位置
	for i := a.size - 1; i >= index; i-- {
		a.array[i+1] = a.array[i]
	}
	//插入新元素
	a.array[index] = element
	a.size++
}

// resize 数组扩容
func (a *MyArray) resize() {
	newArr := make([]int, a.size*2)
	//不能使用append尾插入，需要头插入
	for i := 0; i < a.size; i++ {
		newArr[i] = a.array[i]
	}
	a.array = newArr
}

// Delete 删除元素
func (a *MyArray) Delete(index int) (element int) {
	//判断访问下标是否超出范围
	if index < 0 || index >= a.size {
		panic("超出数组实际元素范围！")
	}
	element = a.array[index]
	//index 之后的元素，往左挪一位
	//TODO:原书错误：i < a.size -1
	for i := index; i <= a.size-1; i++ {
		a.array[i] = a.array[i+1]
	}
	a.size--
	return
}

func (a *MyArray) Output() {
	for i := 0; i < a.size; i++ {
		fmt.Println(a.array[i])
	}
}
