package main

import "fmt"

func fun() (int, int) {
	return 1, 2
}

//注意事项：函数外的每个语句必须以关键字开始 var const func等
// := 不能使用在函数外
// _ 多用于占位，表示忽略值，不会分配内存空间

var bigName string = "大名"

//常量是恒定不变的值，在声明定义的时候必须赋值
const BIGAGE int = 100

//多个常量一起声明
const (
	pi = 3.14
	e  = 2.1
)

//const 同时声明时，如果省略了值则表示和上面一行的值相同
// num1 -- num4 的值都是1
const (
	num1 = 1
	num2
	num3
	num4
)

//iota 枚举
//iota 是go语言的常量计数器，只能在常量的表达式中使用
//iota 在const 关键字出现时将被重置为0，const中每新增一行常量声明将使用iota计数一次
const (
	big1 = iota //0 iota 从 0 开始计数
	big2        //1
	big3        //2
)

//使用 _ 跳过某些值
//打印的值为：0 1 2 4 5
const (
	n1 = iota //0
	n2        //1
	n3        //2
	_         //此处跳过3
	n5        //4
	n6        //5
)

//iota声明中间插队
//打印的值为：0 1 2 10 4 5
const (
	nn1 = iota
	nn2
	nn3
	nn4 = 10
	nn5 = iota
	nn6
)

//定义数量级
// <<< 表示左移操作， 1 << 10 表示 将 1 的二进制向左移动 10 位置 （0001 向左移动10位  （二进制）10000000000 = （十进制的）1024）
// 2 <<< 2 将 2的二进制向左移动 2位 0010 = 1000  即 十进制的 8
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

	//变量的标准声明
	var name string
	var age int

	//变量的批量声明
	var class, address string
	var (
		province string
		city     string
		sort     float32
	)

	//变量的初始化
	var num int = 100

	//变量初始化多个
	var min, max int = 0, 100

	//变量的类型推导，省略类型，直接由右边的值来推导变量的类型来完成初始化
	var count = 200

	//短变量声明
	nameStr := "小明"

	//匿名变量 由_ 下划线替代
	funA, _ := fun()

	/*
	* 整型和浮点型变量的默认值是 0
	* 字符串变量的默认值是 空字符串
	* 布尔型的变量默认值是 false
	* 切片、函数、指针的默认值为 nil
	 */

	name = "cici"
	age = 18
	class = "6班"
	address = "上海市松江区"
	province = "上海"
	city = "上海"
	sort = 740

	fmt.Println(name, age, class, address, province, city, sort, num, min, max, count, nameStr, funA)
	//cici 18 6班 上海市松江区 上海 上海 740 100 0 100 200 小明 1

	//打印常量和外部变量
	fmt.Println(bigName, BIGAGE)

	//打印常量
	fmt.Println(pi, e)
	fmt.Println(num1, num2, num3, num4)

	//打印iota的值, 打印值为：0 1 2
	fmt.Println(big1, big2, big3)

	//打印iota 使用_跳过某一个值，打印结果为：0 1 2 4 5
	fmt.Println(n1, n2, n3, n5, n6)

	//打印iota 中间跳值
	fmt.Println(nn1, nn2, nn3, nn4, nn5, nn6)

	//打印iota的数量级
	fmt.Println(KB, MB, GB, TB, PB)

}
