package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
 * fmt.print 函数直接输出内容
 * fmt.printf 函数支持格式化输出字符串
 * fmt.println 函数会在输出内容的结尾添加一个换行符
 *
 * fmt.Fprint | Fprintf 函数会将内容输出到一个 io.writer 接口类型的变量 w 中，通常用这个函数往文件中写入内容
 * fmt.Sprint | Sprintf 函数会把传入的数据生成并返回一个字符串
 *
 * fmt.Errof 函数根据 format 参数生成格式化字符串并返回一个包含该字符串的错误
 */

/*
 * 通用占位符
 * %v 值的默认格式
 * %+v 类似%v, 但是输出结构体时会添加字段名
 * %#v 值的go语法表示
 * %T 打印值的类型
 * %% 百分号
 *
 *
 * 布尔型
 * %t true | false
 *
 *
 * 整型
 * %b 表示为二进制
 * %c 该值对应的是unicode码值
 * %d 表示为十进制
 * %o 表示八进制
 * %x 表示十六进制，使用 a-f
 * %X 表示为十六进制，使用A-F
 * %U 表示为Unicode格式：U+1234，等价于”U+%04X”
 * %q 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
 *
 *
 * 浮点数与复数
 * %b 无小数部分，二进制指数的科学计数法，如-123456p-78
 * %e 科学计数法，如-1234.456e+78
 * %E 科学计数法，如-1234.456E+78
 * %f 有小数部分但无指数部分，如123.456
 * %F 等价于%f
 * %g 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
 * %G 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
 *
 *
 * 字符串和 []byte
 * %s 直接输出字符串或[]byte
 * %q 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
 * %x 每个字节用两字符十六进制数表示（使用a-f)
 * %X 每个字节用两字符十六进制数表示（使用A-F）
 *
 *
 * 指针
 * %p 表示为十六进制，并加上前导的0x
 *
 *
 * 宽度标识度
 * %f 默认宽度，默认精度
 * %9f 宽度9，默认精度
 * %.2f 默认宽度，精度2
 * %9.2f 宽度9，精度2
 * %9.f 宽度9，精度0
 */

/*
 * fmt 包下有 fmt.Scan, fmt,Scanf, fmt.Scanln
 * 可以通过程序运行过程中从标准输入获取用户的输入
 */

/*
 * 想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可有使用bufio包来实现
 */

func main() {

	//fmt.Fprint
	//fmt.Fprintf
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件报错，err:", err)
	}
	fmt.Fprint(fileObj, "往文件中写入aaaaaaaaa")
	fmt.Fprintf(fileObj, "往文件中写入:%s", "你好")

	//fmt.Sprint | Sprintf 函数会把传入的数据生成并返回一个字符串
	str := fmt.Sprint("你好")
	str1 := fmt.Sprintf("我是：%s", "cici")
	str2 := fmt.Sprintln("我们可以认识吗")
	fmt.Println(str, str1, str2)

	//fmt.Errof 函数根据 format 参数生成格式化字符串并返回一个包含该字符串的错误
	//Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error。
	e := errors.New("报错啦")
	w := fmt.Errorf("报错了吗：%w", e)
	fmt.Println(w)

	//通用占位符
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	//整数型
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)

	//浮点数
	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	//字符串和[]byte
	s := "小王子"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	//指针
	a := 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)

	//宽度表示
	n1 := 12.34
	fmt.Printf("%f\n", n1)
	fmt.Printf("%9f\n", n1)
	fmt.Printf("%.2f\n", n1)
	fmt.Printf("%9.2f\n", n1)
	fmt.Printf("%9.f\n", n1)

	//scan的使用
	// var (
	// 	name  string
	// 	age   int
	// 	isman bool
	// )
	// fmt.Scan(&name, &age, &isman)
	// fmt.Printf("name:%s, age:%d, isman:%t", name, age, isman)

	//scanf的使用
	//指定的格式去读取由空白符分隔的值保存到传递给本函数的参数
	//在终端输入：1:cici 2:18
	// var (
	// 	bigName string
	// 	bigAge  int
	// )
	// fmt.Scanf("1:%s 2:%d", &bigName, &bigAge)
	// fmt.Printf("bigName:%s, bigAge:%d", bigName, bigAge)

	//scanln的使用
	//它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置
	// var (
	// 	minName string
	// 	minAge  int
	// )
	// fmt.Scanln(&minName, &minAge)
	// fmt.Printf("minName:%s, minAge:%d", minName, minAge)

	//从标准输入生成读对象
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	//读到换行
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)

}
