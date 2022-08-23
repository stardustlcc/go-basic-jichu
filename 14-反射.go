package main

import (
	"fmt"
	"reflect"
)

/*
 * go语言中的变量是分为两部分
 * 1）类型信息，预先定义好的元信息
 * 2）值信息，程序运行过程中可动态变化的
 */

/*
 * 反射
 * 反射是指在程序运行期间对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
 *	支持反射的语言可以在程序编译期间将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期间获取类型的反射信息，并且有能力修改它们。
 * Go程序在运行期间使用reflect包访问程序的反射信息
 *
 * 空接口可以存储任意类型的变量，那我们如何知道这个空接口保存的数据是什么呢？ 反射就是在运行时动态的获取一个变量的类型信息和值信息.
 *
 * 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，
 * 并且reflect包提供了 reflect.TypeOf 和 reflect.ValueOf 两个函数来获取任意对象的Value和Type
 * reflect.TypeOf() 可以获得任意值的类型对象（reflect.Type）, 程序可以通过类型对象访问任意值的类型信息
 *
 * 注意：Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
 */

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

/*
 * 反射中关于类型还划分为两种：类型（Type）和种类（Kind）
 * 我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。 举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类
 * 注意：Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
 */
func reflectTypeKind(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

type myInt int64

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

/*
 * 通过反射设置变量的值
 * 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
 */
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("Type:%T val:%v\n", v.Kind(), v.Kind())
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

/*
 * isNil()  报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
 * IsValid() 返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
 * IsNil()常被用于判断指针是否为空；
 * IsValid()常被用于判定返回值是否有效。
 */

/*
 * 结构体反射
 * 任意值通过reflect.TypeOf()获得反射对象信息后，
 * 如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。
 */

func main() {

	var a float32 = 3.14
	reflectType(a) //type:float32
	var b int = 100
	reflectType(b) //reflectType(b)

	//type name | type kind
	var c *float32
	reflectTypeKind(c) //type: kind:ptr // 注意：Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
	var d myInt
	reflectTypeKind(d) //type:myInt kind:int64
	var e rune
	reflectTypeKind(e) //type:int32 kind:int32

	type person struct {
		name string
		age  int
	}

	type book struct {
		title string
	}

	f := person{
		name: "cici",
		age:  19,
	}
	j := book{
		title: "英雄联盟",
	}
	reflectTypeKind(f) //type:person kind:struct
	reflectTypeKind(j) //type:book kind:struct

	var h float32 = 3.14
	var i int64 = 100
	reflectValue(h) //type is float32, value is 3.140000
	reflectValue(i) //type is int64, value is 100

	// 将int类型的原始值转换为reflect.Value类型
	k := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", k) //type c :reflect.Value

	//通过反射设置变量的值
	var g int64 = 100
	reflectSetValue2(&g)
	fmt.Printf("Type:%T val:%v\n", g, g) //Type:int64 val:200
	fmt.Println(g)                       // 200

	// *int类型空指针
	var l *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(l).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	m := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(m).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(m).MethodByName("abc").IsValid())
	// map
	n := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(n).MapIndex(reflect.ValueOf("娜扎")).IsValid())

	//结构体反射案例
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
