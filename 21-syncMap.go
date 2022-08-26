package main

import (
	"fmt"
	"strconv"
	"sync"
)

// var m = make(map[string]int)

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, val int) {
// 	m[key] = val
// }

/*go 语言内置的 map 不是并发安全的，
*如下的案例会报错：fatal error: concurrent map writes
*因为：不能在多个 goroutine中并发对内置的 map 进行读写操作，否则会存在数据竞争问题
*像这种场景下就需要为map加锁来保证并发的安全性。go 中的 sync 提供了开箱即用的并发安全版 map --> sync.map
*sync.Map内置了如下操作
*Store 存储key-val数据
*Load 查询key对应的val
*LoadOrStore 查询或存储key对应的val
*LoadAndDelete 查询并删除key
*Delete 删除key
*Range 对map中的每个key-val依次调用
 */

var m = sync.Map{}

func main() {

	//并发不安全的内置 map 案例
	// wg := sync.WaitGroup{}
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		key := strconv.Itoa(n)
	// 		set(key, n)
	// 		fmt.Printf("k=:%v, v:=%v\n", key, get(key))
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

	//并发安全的sync.Map案例
	wg := sync.WaitGroup{}
	//对m执行10个并发的读写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // 存储key-value
			value, _ := m.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
