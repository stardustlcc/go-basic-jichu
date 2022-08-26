package main

import (
	"fmt"
	"sync"
)

var (
	x  int64
	wg sync.WaitGroup
	m  sync.Mutex //增加互斥锁
)

/*
 * 互斥锁
 * 互斥锁是一种常用的控制共享资源访问的方法，可以保证同一时间只有一个goroutine可以访问共享资源，go语言使用 sync 包中提供的 Mutex 类型来实现互斥锁。
 *
 * 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁，当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
 */
func add() {
	for i := 0; i < 5000; i++ {
		//修改x前加锁
		m.Lock()
		x = x + 1
		//修改后释放锁
		m.Unlock()
	}
	wg.Done()
}

func main() {

	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	//此时 x 的值每次执行都产生了不一样的值，因为产生了竞争，有被重写的风险。需要使用互斥锁解决问题
	fmt.Println(x)

}
