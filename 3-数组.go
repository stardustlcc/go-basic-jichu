package main

import "fmt"

/*
 * 数组是同一种数据类型元素的集合，在go语言中，数组从声明就确定，使用时可以修改数组成员，但是数组大小不可修改
 */
func main() {

	//数组的长度一旦定义则不可修改，长度不同的两个数组是不同的类型
	var arr [3]string

	arr[0] = "你好"
	arr[1] = "我是"
	arr[2] = "你的朋友"

	fmt.Println(arr)

	//方法一：数组的初始化
	var arr1 [5]int                     //数组会初始化为int类型的零值
	var arr2 = [4]string{"a", "b", "c"} //使用指定的初始值完成初始化
	fmt.Println(arr1)
	fmt.Println(arr2)

	//方法二：使用编辑器帮我们自行推断数组的长度
	var arr3 = [...]int{1, 2, 3, 4, 5, 6}
	var arr4 = [...]string{"上海", "北京"}
	fmt.Println(arr3)
	fmt.Println(arr4)

	//方法三：使用指定索引值的方式来初始化数组
	arr5 := [...]string{1: "aaa", 3: "bbb"}
	fmt.Println(arr5)

	//数组的遍历
	for index, value := range arr4 {
		fmt.Println(index, value)
	}

	//多维数组
	arr6 := [3][2]string{
		{"aaa", "aaa"},
		{"bbb", "bbb"},
		{"ccc", "ccc"},
	}
	fmt.Println(arr6)

	//多维数组的遍历
	for _, value := range arr6 {
		for i, v := range value {
			fmt.Println(i, v)
		}
	}

	//多维数组只有第一层可以使用 ... 让编译器推导数组的长度
	arr7 := [...][2]string{
		{"aaa", "aaa"},
		{"bbb", "bbb"},
	}
	fmt.Println(arr7)

	//数组是值类型，赋值和传参都会赋值整个数组，因此改变副本的值，不会改变本身的值
	arr8 := [3]int{1, 2, 3}
	fmt.Println(arr8)
	funcTest(arr8)
	fmt.Println(arr8)

}

func funcTest(arr [3]int) {
	arr[1] = 100
	fmt.Println(arr)
	return
}
