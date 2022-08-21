package main

import (
	"fmt"
)

/*
 * map是一种无序的基于key-val的数据结构，map是引用类型，必须初始化才能使用
 */
func main() {

	//使用make初始化一个map变量
	var mapOne = make(map[string]string, 3)
	mapOne["name"] = "cici"
	mapOne["age"] = "18"
	fmt.Println(mapOne)

	mapTwo := make(map[int]string, 10)
	mapTwo[0] = "aaa"
	mapTwo[1] = "bbb"
	fmt.Println(mapTwo)

	userInfo := map[string]string{
		"name": "cici",
		"age":  "18",
	}
	fmt.Println(userInfo)

	//判断某个键是否存
	val, ok := userInfo["name"]
	if ok {
		fmt.Println(val)
	}

	//map的遍历
	for k, v := range userInfo {
		fmt.Println(k, v)
	}

	//使用delete()删除键值对
	delete(userInfo, "name")
	fmt.Println(userInfo)

	//按照指定顺序遍历map

	//元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "cici"
	mapSlice[0]["age"] = "19"
	fmt.Println(mapSlice) //[map[age:19 name:cici] map[] map[]]

	//值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	sliceMap["test"] = make([]string, 3)
	sliceMap["test"][0] = "cici"
	fmt.Println(sliceMap)

}
