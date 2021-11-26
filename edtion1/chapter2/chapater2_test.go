package chapter2

import (
	"ProgrammerXiaohuiGolang/edtion1/chapter2/part1"
	"ProgrammerXiaohuiGolang/edtion1/chapter2/part2"
	"ProgrammerXiaohuiGolang/edtion1/chapter2/part3"
	"fmt"
	"testing"
)

func TestNewMyArray(t *testing.T) {
	myArray := part1.NewMyArray(4)
	myArray.Insert(0,0)
	myArray.Insert(1,1)
	myArray.Insert(2,2)
	myArray.Insert(3,3)
	myArray.Insert(4,4)
	myArray.Insert(5,5)
	myArray.Delete(3)
	fmt.Println(myArray.Array)
	myArray.Output()
}

func TestNewMyLinkedLink(t *testing.T)  {
	myLinkedLink := part2.NewMyLinkedLink()
	myLinkedLink.Insert(0,3)
	myLinkedLink.Insert(0,1)
	myLinkedLink.Insert(2,4)
	myLinkedLink.Insert(3,5)
	myLinkedLink.Insert(1,2)
	myLinkedLink.Remove(3)
	myLinkedLink.Output()
}

func TestNewQueue(t *testing.T)  {
	myQueue := part3.NewMyQueue(6)
	myQueue.EnQueue(1)
	myQueue.EnQueue(2)
	myQueue.EnQueue(3)
	myQueue.EnQueue(4)
	myQueue.DeQueue()
	myQueue.DeQueue()
	myQueue.EnQueue(4)
	myQueue.EnQueue(5)
	myQueue.EnQueue(6)
	//myQueue.EnQueue(7)
	myQueue.Output()
}
