package main

import (
	"fmt"
	"sync"
)

type Student struct {
	name string
}

var (
	once sync.Once
	stu  *Student
)

func myTest() *Student {
	//只会执行一次
	//适用于单例模式
	once.Do(func() {
		stu = new(Student)
		stu.name = "cici"
		fmt.Println("new student cici")
	})
	//会执行多次
	fmt.Println("get cici student")
	return stu
}

func main() {

	for i := 0; i < 10; i++ {
		_ = myTest()
	}
}
