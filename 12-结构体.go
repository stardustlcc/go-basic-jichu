package main

import (
	"fmt"
	"unsafe"
)

/*
 * 自定义类型
 * 自定义类型是定义了一个全新的类型，我们可以基于内置的基本类型定义，也可以通过struct定义
 */

/*
 * 类型别名
 * 类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名
 *
 * 类型定义和类型别名的区别
 * 表示main包下定义的NewInt类型。b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型
 */

/*
 * 结构体
 * 使用type和struct关键字来定义结构体
 * 结构体实例化
 * var 结构体实例 结构体类型
 *
 * 结构体实例化
 * 1）使用new关键字对结构体进行实例化，得到的是结构体的地址
 * 2）使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
 * 结构体初始化
 * 没有初始化的结构体，其成员变量都是对应其类型的零值
 */

/*
 * 构造函数
 * go语言中没有构造函数，我们可以自己实现，因为 struct 是值类型，如果结构体比较复杂，值拷贝性能开销会比较大，所以够着函数返回的是结构体指针类型
 */

//结构体
//结构体用来描述一组值
type person struct {
	name, city string
	age        int8
}

func main() {

	//自定义类型
	type myInt int
	//通过Type关键字定义，myint 就是一种新的类型，它具有int的特性

	var age myInt
	age = 18
	fmt.Printf("Type:%T val:%#v\n", age, age) //Type:main.myInt val:18\\

	//类型别名
	//type byte = uint8
	//type rune = int32

	//结构体实例化
	var p1 person
	p1.age = 18
	p1.name = "cici"
	p1.city = "安徽"
	fmt.Println(p1)

	//匿名结构体
	t0 := struct {
		name string
	}{
		name: "cici",
	}
	fmt.Printf("%#v\n", t0) //struct { name string }{name:"cici"}

	//匿名字段
	type user struct {
		string
		Age int
	}
	t1 := user{string: "cici", Age: 18}
	fmt.Printf("%#v\n", t1) //main.user{string:"cici", Age:18}

	//创建指针类型结构体
	//我们可以通过new 对关键字结构体进行实例化，得到的是结构体的地址
	var p2 = new(person)
	p2.name = "cici"
	p2.age = 19
	p2.city = "上海"
	fmt.Printf("Type:%T, Val:%#v\n", p2, p2) //Type:*main.person, Val:&main.person{name:"cici", city:"上海", age:19}

	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	var p3 = &person{}
	p3.name = "cici"
	p3.city = "北京"
	p3.age = 20
	fmt.Printf("Type:%T, Val:%#v\n", p3, p3) //Type:*main.person, Val:&main.person{name:"cici", city:"北京", age:20}

	//结构体初始化
	var p4 person
	fmt.Printf("Type:%T, Val:%#v\n", p4, p4) //Type:main.person, Val:main.person{name:"", city:"", age:0}
	//使用键值对初始化
	p5 := person{
		name: "cici",
		city: "武汉",
		age:  18,
	}
	fmt.Printf("Type:%T, Val:%#v\n", p5, p5) //Type:main.person, Val:main.person{name:"cici", city:"武汉", age:18}
	//使用结构体指针进行键值对初始化
	p6 := &person{
		name: "ciic",
		city: "四川",
		age:  19,
	}
	fmt.Printf("Type:%T, Val:%#v\n", p6, p6) //Type:*main.person, Val:&main.person{name:"ciic", city:"四川", age:19}
	//使用值的列表初始化，可不写键，直接写值
	p7 := person{
		"cici",
		"安徽",
		18,
	}
	fmt.Printf("Type:%T, Val:%#v\n", p7, p7) //Type:main.person, Val:main.person{name:"cici", city:"安徽", age:18}

	//空结构体是不占用空间的。
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0

	//调用构造函数
	p8 := newPerson("cici", "广州", 19)
	fmt.Printf("Type:%T, Val:%#v\n", p8, p8) //Type:*main.person, Val:&main.person{name:"cici", city:"广州", age:19}
	p8.Dream()

	//指针类型的接收者
	//指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的this或者self
	p8.setAge(22)
	fmt.Printf("Type:%T, Val:%#v\n", p8, p8) //Type:*main.person, Val:&main.person{name:"cici", city:"广州", age:22}

	//什么时候应该使用指针类型接收者
	//需要修改接收者中的值
	//接收者是拷贝代价比较大的大对象
	//保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

}

//构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//person 的方法 值类型的接受者
func (p person) Dream() {
	fmt.Println(p.name, "开始做梦了")
}

//person 的方法 指针类型的接受者
func (p *person) setAge(age int8) {
	p.age = age
}
