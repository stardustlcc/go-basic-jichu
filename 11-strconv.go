package main

import (
	"fmt"
	"strconv"
)

/*
 * strconv包实现了基本数据类型与字符串表示的转换
 * 主要有：Atoi() Itoa() parse系列，formar系列，append系列
 * Atoi() 函数用于将字符串类型的整数转换为int类型
 * Itoa() 函数用于将int类型数据转换为对应的字符串表示
 */

/*
 * parse系列函数用于转换字符串为给定类型的值
 * parseBool()  parseFloat() parseInt() parseUint()
 */

/*
 * format系列函数实现了将给定类型数据格式化为string类型数据的功能
 * FormatBool() 返回”true”或”false”。
 * FormatInt()
 * FormatUint()
 * FormatFloat()
 */
func main() {

	//string 和 int 类型转换
	str1 := "100"
	int1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("cant convet to int")
	} else {
		fmt.Printf("type:%T val:%#v\n", int1, int1) //type:int val:100
	}

	//Itoa()函数用于将int类型数据转换为对应的字符串表示
	int2 := 200
	str2 := strconv.Itoa(int2)
	fmt.Printf("Type:%T val:%#v\n", str2, str2) //Type:string val:"200"

	//parseBool() 它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误
	str3 := "T"
	bool1, err := strconv.ParseBool(str3)
	if err != nil {
		fmt.Println("类型转换失败", err)
	}
	fmt.Printf("Type:%T Val:%#v\n", bool1, bool1) //Type:bool Val:true

	//parseInt() 返回字符串表示的整数值，接受正负号
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	str4 := "127"
	int3, err := strconv.ParseInt(str4, 10, 8) //这里的16表示结果最大值不超过int8即127
	if err != nil {
		fmt.Println("类型转换失败", err)
	}
	fmt.Printf("Type:%T Val:%#v\n", int3, int3)

	//ParseUnit()
	//ParseUint类似ParseInt但不接受正负号，用于无符号整型。

	//ParseFloat()
	//解析一个表示浮点数的字符串并返回其值。
	//bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
	str5 := "1.23"
	flot1, _ := strconv.ParseFloat(str5, 64)
	fmt.Printf("Type:%T Val:%#v\n", flot1, flot1) //Type:float64 Val:1.23

	//format 系列
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Printf("Type:%T Val:%#v\n", s1, s1) //Type:string Val:"true"
	fmt.Printf("Type:%T Val:%#v\n", s2, s2) //Type:string Val:"3.1415E+00"
	fmt.Printf("Type:%T Val:%#v\n", s3, s3) //Type:string Val:"-2"
	fmt.Printf("Type:%T Val:%#v\n", s4, s4) //Type:string Val:"2"
}
