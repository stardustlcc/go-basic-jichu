package main

import (
	"fmt"
	"math"
)

/*
 * go 语言中 除了基本的 整型，浮点型，布尔型，字符串，还有数组，切片，结构体，函数，map，通道
 */

/*
 * 整型分为两类：有符号整型，无符号整型
 * 有符号整型：int8  int16  int32  int64
 * 无符号整型：uint8 uint16 uint32 uint64
 *
 * 注意：uint8 = byte 型
 *
 * uint8  无符号8位整型（0-255）
 * uint16 无符号16位（0-65535）
 * uint32 无符号32位 （0-4294967295）
 * uint64 无符号64位 （0-18446744073709551615）
 *
 * int8  有符号8位整型 （-128 -- 127）
 * int16 有符号16位整型 （-32768 -- 32767）
 * int32 有符号32位整型 （-2147483648 -- 2147483647）
 * int64 有符号64位整型 （-9223372036854775808 -- 9223372036854775807）
 *
 * 注意：特殊整型
 * uint    32位操作系统上就是 uint32 , 64位操作系统上就是 uint64
 * int     32位操作系统上就是 int32 , 64位操作系统就是 int64
 * uintptr 无符号整型，用于存放一个指针
 *
 * 注意： uint 和 int 在不同平台上的差异。在使用len()函数时，为了保证文件结构不受变异平台字节长度的影响，不要使用 int 和 uint
 */

/*
 * 浮点型有两种：float32 float64 。
 *            float32的最大范围约 3.4e38 可以使用常量定义 math.MaxFloat32
 * 		   float64的最大范围约 1.8e308 可以使用常量定义 math.MaxFloat64
 * 打印浮点型，可以使用 fmt 包 配合动词 %f
 */

/*
 * 复数：complex64 和 complex128
 * 复数有实部和虚部 complex64的实部和虚部为32  complex128的实部和虚部为64
 */

/*
 * 布尔型 bool (true | false) 两个值
 * 布尔型默认值为 false
 * go语言中不允许将整型强制转换为布尔型
 * 布尔型无法参与数值运算，也无法与其他类型进行转换
 */

/*
 * byte 和 rune类型
 * 组成每个字符串的元素叫做“字符” 可以通过遍历或者单个获取字符串元素获得字符。字符用单引号包裹起来
 * go语言的字符有以下两种
 *  1）uint8类型，或者叫byte型，代表一个ASCII码字符
 *  2）rune类型，代表一个UTF-8字符
 *  当处理中文，或其他复合字符时，则需要使用到 rune 类型，rune 类型实际是一个 int32
 *
 */

/*
 * 遍历字符串
 * hello你好 输出结果：104(h)101(e)108(l)108(l)111(o)228(ä)189(½)160( )229(å)165(¥)189(½)
 * 说明：因为UTF8编码下的一个中文汉字由3-4个字节组成，所以不能使用字节遍历一个包含中文的字符串！
 * 当处理中文，或其他复合字符时，则需要使用到 rune 类型，rune 类型实际是一个 int32
 * go 使用了特殊的 rune 类型来处理 unicode，让基于unicode的文本处理更为方便，也可以使用byte型进行默认字符串处理，性能和扩展都有照顾
 */

/*
 * 修改字符串
 * 字符串底层是一个byte数组，所以可以使用 []byte类型相互转换，字符串是补鞥呢修改的，字符串是由byte字节组成，所以字符串的长度是 byte字节的长度，
 * rune类型采用 utf8字符，一个rune字符由一个或多个byte组成
 * 将字符串转换成 []rune 或 []byte 完成后在转换成 string。无论哪种转换，都会重新分配内存，并复制字节数组
 */

/*
 * 类型转换
 * T(表达式)
 */

func main() {

	//浮点型案例输出
	var num float32 = 3.14
	fmt.Printf("%f\n", num)    //3.140000
	fmt.Printf("%.2f\n", 3.14) //3.14

	//复数案例
	var c1 complex64
	var c2 complex128
	c1 = 1 + 2i
	c2 = 1 + 3i
	fmt.Println(c1, c2) //(1+2i) (1+3i)

	//字符的案例
	var char = 'a'
	fmt.Println(char)

	//遍历字符 temp1 byte
	//输出结果：104(h)101(e)108(l)108(l)111(o)228(ä)189(½)160( )229(å)165(¥)189(½)
	//说明：因为UTF8编码下的一个中文汉字由3-4个字节组成，所以不能使用字节遍历一个包含中文的字符串！
	//go 使用了特殊的 rune 类型来处理 unicode，让基于unicode的文本处理更为方便，也可以使用byte型进行默认字符串处理，性能和扩展都有照顾
	var str = "hello你好"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v(%c)", str[i], str[i])
	}
	fmt.Println()

	//遍历字符串 temp2 rune
	for _, r := range str {
		fmt.Printf("%v(%c)", r, r)
	}

	//修改字符串
	//字符串底层是一个byte数组，所以可以使用 []byte类型相互转换，字符串是补鞥呢修改的，字符串是由byte字节组成，所以字符串的长度是 byte字节的长度，
	//rune类型采用 utf8字符，一个rune字符由一个或多个byte组成
	//将字符串转换成 []rune 或 []byte 完成后在转换成 string。无论哪种转换，都会重新分配内存，并复制字节数组
	byteStr := []byte(str)
	byteStr[0] = 'A'
	//byteStr[5] = '我' 不可修改中文
	fmt.Println(string(byteStr))

	byteStrRune := []rune(str)
	byteStrRune[5] = '我'
	fmt.Println(string(byteStrRune))

	//类型转换
	//math.Sqrt()接收的参数是float64类型，需要强制转换
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * b)))
	fmt.Println(c)

}
