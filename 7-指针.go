package main

import "fmt"

/*
 * 每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置，
 * 使用 & 字符对变量进行取地址操作
 * 使用 * 字符对指针变量进行取值操作
 */

/*
 * go语言对于引用类型的变量，我们不仅要声明它，还要为它分配内存空间，不然没有办法存储。对于值类型的变量不需要分配内存空间，是因为他们在声明的时候已经默认分配好了内存空间
 * 指针如果只是声明了指针变量，并没有初始化，则需要使用 new or make 进行初始化申请内存空间，方可使用，否则会报错！！
 * new函数不常用，使用new得到的是一个类型的指针，并且该指针对应的值为该类型的零值
 * make函数也是内存分配的，区别于new,它只用于slice map 和 channel的内存创建，返回的类型就是这三个类型本身，而不是他们的指针类型。因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
 */

func main() {

	num := 100
	newNum := &num
	fmt.Println(newNum) // 0xc000014098 是 num的地址

	//newNum是指针变量
	fmt.Println(*newNum) //100 对着指针变量取值

	//new() 的案例
	//说明：new函数不常用，使用new得到的是一个类型的指针，并且该指针对应的值为该类型的零值
	a := new(int)
	*a = 100
	fmt.Println(*a)
	b := &a
	fmt.Println(b)   //0xc0000ac020
	fmt.Println(**b) //100

	//new() 函数为指针变量初始化内存
	var c *int
	c = new(int)
	*c = 200
	fmt.Println(*c) //200

	//make函数的使用,只能用于 slice,map,channel
	d := make(map[int]string, 1)
	d[1] = "cici"
	fmt.Println(d)
}
