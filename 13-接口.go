package main

import (
	"fmt"
)

/*
 * 接口是一种类型，一种抽象类型。
 * 接口类型就像是一种约定，概括了一种类型应该具备哪些方法，在go语言中提倡使用面向接口的编程的方式实现解耦
 *
 * 接口是一种由程序员定义的类型，一个接口类型就是一组方法的集合，它规定了要实现的所有方法
 * 相比较结构体类型，当我们使用接口类型说明相比较它是什么更关系它能做什么
 */

//定义一个包含Write方法的Writer接口
//接口就是规定了一个需要实现的方法列表
//在go语言中一个类型只要实现了接口中规定的所有方法，就称它实现了这个接口
type Writer interface {
	Write([]byte) error
}

//定义一个Singer接口类型，它包含一个Sing方法
type Singer interface {
	Sing()
}

type Person struct {
	name string
	age  int
}

//我们称Person 实现了 Singer 接口
func (p Person) Sing() {
	fmt.Println("人类会唱歌")
}

/*
 * 一个类型实现多个接口，是没有问题的，每个接口彼此独立，不知道对方的实现
 */
type Sayer interface {
	Say()
}
type Mover interface {
	Move()
}

//定义一个狗类型
type Dog struct {
	Name string
}

func (d Dog) Say() {
	fmt.Println("汪汪汪")
}

func (d Dog) Move() {
	fmt.Println("小狗开始走路了")
}

/*
 * 多种类型实现同一接口，例如不仅狗狗会动，骑车也会动
 */
type Car struct {
	Brand string
}

func (c Car) Move() {
	fmt.Println("小汽车在移动")
}

/*
 * 一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型嵌套入其他类型或者结构体来实现
 */
//洗衣机接口
type WashingMachine interface {
	wash() //洗衣服
	dry()  //烘干
}

//烘干机
type dryer struct{}

//海尔洗衣机
type haier struct {
	dryer //嵌入烘干机
}

func (h haier) wash() {
	fmt.Println("洗衣服")
}

func (d dryer) dry() {
	fmt.Println("烘干衣服")
}

/*
 * 接口组合
 * 接口与接口之间可以通过相互嵌套成新的接口类型
 */
type Reader interface {
	Read()
}
type Writers interface {
	Write()
}
type Closer interface {
	Close()
}

//组合接口
//这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。
type ReadWriter interface {
	Reader
	Writers
}
type ReadCloser interface {
	Reader
	Closer
}
type WriteClose interface {
	Writers
	Closer
}

/*
 * 接口也可以做结构体的一个字段
 */
type Interface interface {
	Len(i, j int) bool
}
type reverse struct {
	Interface
}

//结构体实现了接口的方法:并且重写了参数交换位置实现反转
func (r reverse) Len(i, j int) bool {
	return r.Interface.Len(j, i)
}

/*
 * 空接口
 * 空接口是指没有定义任何方法的接口类型，因此任何类型都可以视为实现了空接口，也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值
 * 通常使用空接口类型时不必使用 type 关键字声明可以这样用：var x interface{}  // 声明一个空接口类型变量x
 */

//不包含任何方法的空接口类型
type Any interface{}

//不适用type定义空接口
//var SomeAny interface{}

//定义一个结构体
type Dogs struct{}

/*
 * 一个接口类型 有一个方法 (接口有：接口动态类型、接口动态值)
 * 1）创建一个接口类型的变量 m
 * var m Mover
 * 此时 m == nil 即判断此时的接口值是nil
 * 注意：我们不能对一个空 nil 接口值调用任何方法，否则会panic
 * 2) 接下来将一个结构体指针赋值给一个接口变量 m
 * m = &Dog{Name: "旺财"}
 * 此时 m 的动态类型是*Dog 动态值为结构体变量的拷贝 Dog{旺财}
 * 此时 m != nil
 * 如果 &Dog 里面没有值 Dog{}, m 仍旧是 不等于 nil 。 因为它只是动态值的部分为nil,而动态类型部分保存着对应的类型
 *
 * 注意：接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等才相等
 * 3）如果接口值保存的动态类型相同，值也相同，即自己跟自己比较，如果这个动态类型不支持相互比较（比如切片）那么他们相互比较会引发panic
 * var z interface{} = []int{1, 2, 3}
 * fmt.Println(z == z) // panic: runtime error: comparing uncomparable type []int
 */

/*
 * 类型断言
 * 接口值，尤其是空接口的接口值，可能赋值为任意类型的值，那如何从接口值获取其存储的具体数据呢？？
 * 1）可以借助fmt 包格式打印接口值的动态类型
 * 而 fmt 包内部其实是使用反射的机制在程序运行时获取到动态类型的名称
 * 2) 而想要从接口值中获取到对应的实际值需要使用类型断言
 * 语法：x.(T)
 * x 接口类型的变量
 * T 表示断言，x可能是的类型
 * 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
 */

func main() {

	var d Dog
	d.Name = "小黑"
	d.Say()

	var obj Mover
	obj = Dog{Name: "小黑"}
	obj.Move()

	obj = Car{Brand: "奥迪"}
	obj.Move()

	var h haier
	h.dry()
	h.wash()

	//空接口类型
	var a Any
	a = "你好"
	fmt.Printf("type:%T value:%v\n", a, a) //type:string value:你好

	a = 200
	fmt.Printf("type:%T value:%v\n", a, a) //type:int value:200

	a = Dogs{}
	fmt.Printf("type:%T value:%v\n", a, a) //type:main.Dogs value:{}

	//声明一个空接口
	//空接口作为参数
	var SomeAny interface{}
	SomeAny = Dogs{}
	fmt.Printf("type:%T value:%v\n", SomeAny, SomeAny) //type:main.Dogs value:{}
	show(SomeAny)                                      //空接口作为参数 type:main.Dogs value:{}

	SomeAny = true
	fmt.Printf("type:%T value:%v\n", SomeAny, SomeAny) //type:bool value:true
	show(SomeAny)                                      //空接口作为参数 type:bool value:true

	//空接口作为map的值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "cici"
	studentInfo["age"] = 18
	studentInfo["isMax"] = true
	fmt.Println(studentInfo) //map[age:18 isMax:true name:cici]

	//接口类型和接口值的比较
	//var z interface{} = []int{1, 2, 3}
	//fmt.Println(z == z) //panic: runtime error: comparing uncomparable type []int 切片不支持比较哦

	//类型断言
	//1)通过fmt 包 的格式化打印获取接口值的动态类型
	var s Any
	s = &Dogs{}
	fmt.Printf("type:%T value:%v\n", s, s) //type:*main.Dogs value:&{}

	s = new(Dog)
	fmt.Printf("type:%T value:%v\n", s, s) //type:*main.Dog value:&{}

	var testType interface{}
	testType = 100
	getType(testType) //是整型 100

	testType = "你好 我是xxx" //是字符串 你好，我是xxx
	getType(testType)

	testType = false
	getType(testType) //是布尔型 false

	testType = Dogs{}
	getType(testType)

}

//空接口的应用
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

//类型断言
func getType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Println("是字符串", v)
	case int:
		fmt.Println("是整型", v)
	case bool:
		fmt.Println("是布尔型", v)
	default:
		fmt.Println("不知道的类型")
	}
}
