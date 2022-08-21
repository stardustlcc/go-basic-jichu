package main

import "fmt"

/*
 * 数组的缺点
 * 因为数组的长度固定，并且数组长度属于类型的一部分，所以数组有很多的局限性，比如传参的形参，一旦确定了类型，则不够灵活
 *
 * 切片是拥有相同类型元素的可变长度的序列，它是基于数组类型的一层封装，非常灵活，支持自动扩容
 * 切片是引用类型，内部结构包含：地址，长度，和容量
 * len() 求切片的长度
 * cap() 求切片的容量
 */

/*
 * 切片的底层就是一个数组，所以可以基于数组通过切片表达式得到切片（左包括，右不包括）
 */

/*
 * append 可以为切片动态添加元素，可以添加一个元素，也可以添加多个元素，也可以添加另一个切片中的元素，追加到后面
 */

/*
 * 使用copy函数复制切片
 * 为什么需要拷贝呢，而不是直接赋值切片呢？
 * 因为切片是引用类型，所以一个切片赋值给另一个变量，那么两个变量其实同时指向了同一块内存地址。
 * 如果我们不想让它们复用一个内存地址就需要使用到copy()
 */

/*
 * 从切片中删除元素
 * go语言中并没有删除切片元素的专用方法，可以使用切片本身的特性来删除元素
 */

func main() {

	//声明一个int类型的切片
	var a []int

	//声明一个字符串切片并初始化
	var b = []string{}

	//声明一个布尔切片并初始化
	var c = []bool{false, true}

	//声明一个布尔切片并初始化
	var d = []bool{false, true}

	fmt.Println(a) //[]
	fmt.Println(b) //[]
	fmt.Println(c) //[false true]
	fmt.Println(d) //[false true]

	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	//fmt.Println(c == d) 切片是引用类型，不支持直接比较，只能和nil比较

	//1)基于数组创建切片
	arr := [6]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println(slice)                                    //[2 3 4]
	fmt.Printf("len(%v) cap(%v)", len(slice), cap(slice)) //len(3) cap(5)

	arr1 := [5]int{1, 2, 3, 4, 5}
	t := arr1[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t)) //len(3) cap(5)t:[2 3] len(t):2 cap(t):4

	//2)使用make()函数构造切片
	slice1 := make([]int, 2, 10) //2是长度，10 是容量
	fmt.Println(slice1)

	//判断切片是否是空
	//检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
	//切片之间不能比较，切片只能和nil进行比较。一个nil值的切片底层没有数组，没有被初始化内存空间
	if len(slice1) == 0 {
		fmt.Println("切片为空")
	}
	var slice2 []int
	if slice2 == nil {
		fmt.Println("切片没有初始化")
	}

	//切片的赋值
	slice3 := make([]int, 3)
	slice2 = slice3
	slice2[1] = 100
	fmt.Println(slice2, slice3) //[0 100 0] [0 100 0]

	//切片的遍历
	slice4 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(slice4); i++ {
		fmt.Println(slice4[i])
	}

	for index, val := range slice4 {
		fmt.Println(index, val)
	}

	//append方法为切片添加元素
	//通过var 声明的零值切片可以 直接使用 append，不需要初始化，因为append底层会自动扩容，申请内存，所以每次都要重新用变量接值
	var slice5 []int
	slice5 = append(slice5, 1)
	slice5 = append(slice5, 2)
	slice5 = append(slice5, 3)
	slice5 = append(slice5, 4, 5, 6)
	fmt.Println(slice5) //[1 2 3 4 5 6]

	//copy函数的使用
	slice6 := []int{1, 2, 3}
	slice7 := slice6
	fmt.Println(slice7) //[1 2 3]
	slice6[1] = 100
	fmt.Println(slice7) //[1 100 3]

	slice8 := make([]int, 4)
	copy(slice8, slice6)
	fmt.Println(slice8) //[1 100 3 0]
	slice6[2] = 200
	fmt.Println(slice8) //[1 100 3 0]

	//从切片中删除元素
	slice9 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//删除索引为3的元素
	slice9 = append(slice9[:3], slice9[4:]...)
	fmt.Println(slice9) //[1 2 3 5 6 7 8 9]

}
