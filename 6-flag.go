package main

import (
	"flag"
	"fmt"
)

/*
 * go内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单
 * os.Args 是一个存储命令行参数的字符串切片。它可以简单的获取命令行参数
 */

/*
 * flag 包的基本使用
 * flag 包支持的命令行参数类型有 bool，int, int64, uint, uint64,float,float64,string,duration
 *
 * flag.Type 其中 Type 由 string\int\bool\duration
 * 格式：flag.Type(flage名，默认值，帮助信息) *Type
 * 返回的结果是指针类型
 * 需要使用 flag.Parse()来对命令行参数进行解析
 *
 *
 * flag.TypeVar()
 * 格式：flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
 *
 * 支持命令行参数的格式有以下几种
 * -flag xxx （使用空格，一个-符号）
 * --flag xxx （使用空格，两个-符号）
 * -flag=xxx （使用等号，一个-符号）
 * --flag=xxx （使用等号，两个-符号）
 * 注意：布尔类型的参数必须使用等号的方式指定
 *
 */

func main() {

	//os.args 的操作使用 go run 6-flag.go aaa bbb ccc
	// if len(os.Args) > 0 {
	// 	for index, val := range os.Args {
	// 		fmt.Println(index, val)
	// 	}
	// }

	//flag.Type的案例，返回的结果是指针类型
	//操作使用：go run 6-flag.go -name aaa -age 10
	//name := flag.String("name", "cici", "请输入你的姓名")
	//age := flag.Int("age", 18, "请输入你的年龄")

	//使用parse来解析
	//flag.Parse()

	//fmt.Println(*name, *age)

	//flag.TypeVar()的案例
	var name1 string
	var age1 int
	flag.StringVar(&name1, "name", "cici", "姓名")
	flag.IntVar(&age1, "age", 18, "年龄")
	flag.Parse()
	fmt.Println(name1, age1)

}
