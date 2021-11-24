package chapter2

import (
	"ProgrammerXiaohuiGolang/edtion1/chapter2/part1"
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